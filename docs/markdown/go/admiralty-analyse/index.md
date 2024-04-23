# Admiralty - 基于 Virtual Kubelet 实现的跨集群调度平面

## 前言

目前在车企负责一个包括出车镜像构建、单元测试、仿真、感知评测的通过k8s operator实现的任务流，基于供应链与成本相关的原因，需要加入多架构镜像构建（分别在amd64与arm64的集群上构建出车镜像并通过manifest的方式组织出一个多架构镜像），之前看到了Admiralty可以实现低成本的跨集群任务调度，但是并没有细究其内部实现，于是希望好好学习一番。

## 工作流程

跨集群调度功能需要在命名空间级别启用标签`multiclusterscheduler=enabled`，在pod级别启用标签`multicluster.ammanty.io/elect=""`，在被标记命名空间中含有标签的pod被称为源pod。

![](./scheduling.svg)

跨集群调度的工作通过四个组件共同实现：

> 1.针对变异pod的准入程序
> 
> 2.代理调度器
> 
> 3.伴侣pod控制器
> 
> 4.候选调度器

对于一个给定的变异pod，回调器和代理调度器在源集群运行，伴侣pod控制器和候选调度器在源集群所对应的目标集群中工作。

> 代理和候选调度器使用[Scheduler Framework插件](https://kubernetes.io/docs/concepts/scheduling-eviction/scheduling-framework/)扩展了标准的Kubernetes调度器。因此，它们实现了所有标准的Kubernetes调度功能。

### 代理pod

变异pod准入程序将源pod转换成代理pod:

> 1.原始清单保存为注释，以便过后用于候选pod
> 
> 2.调度约束（节点选择器、关联性、容忍度和拓扑扩展约束）被剥离和替换，以选择和容忍代表目标集群的虚拟节点
> 
> 3.调度方式选择为代理调度器`admiralty-proxy`

代理调度程序把代理pod调度并绑定到代表目标集群的虚拟节点。在代理pod的调度周期中，代理调度器在每个目标集群中(在同一个名称空间中)创建一个候选pod伴侣，并等待该集群中的候选调度器为相应的候选pod预留一个节点，或者不预留。

### 伴侣pod

一种crd（自定义资源），它们像pod一样有spec和status，它们被用作代理pod和候选pod间的中间对象。代理调度器会为了创建候选pod伴侣做以下事情:

> 1.使用源pod的spec（在代理pod中被保存为注释）
>
> 2.调度方式选择为代理调度器`admiralty-candidate`

每个目标集群的pod伴侣控制器使用相同的规范创建一个候选pod。稍后，控制器还会根据候选pod的状态更新pod伴侣的状态和注释。如果候选pod被直接删除(不是因为代理pod被删除) ，例如，如果它被驱逐或抢占，控制器将确保它被替换，即使源集群和目标集群之间的连接被中断。

> 在目标集群中，源集群标识只允许与pod伴侣交互，而不能直接与pod交互，因此它们不能读/写非候选pod。这遵循最小特权原则。

### 候选pod

每个目标集群的候选调度程序将候选pod调度到常规节点。

在候选Pod的调度周期结束时，如果一个节点已经被预留（已经过滤、计分，但还未绑定），在此节点后运行的插件会用`multicluster.admiralty.io/is-reserved=true`对Pod伴侣进行注解。这是为了通知位于源集群中的代理调度器，在其过滤插件等待中，表示目标集群的虚拟节点是代理Pod的潜在匹配项。

如果候选 pod 不可调度，代理调度程序将从pod伴侣的状态(反映候选pod的状态)看到它，并停止等待，拒绝虚拟节点。

在前一种情况下，如果候选pod是可调度的，那么在插件的绑定周期开始时，候选调度程序将等待代理调度程序选择候选pod作为委托pod，委托pod只能有一个。

### 委托pod

如果一个虚拟节点通过了代理调度器的过滤器插件，它仍然必须通过其他标准的Kubernetes过滤器(例如，当它的集群被升级时，虚拟节点是否被封锁?).当多个虚拟节点通过所有过滤器时(更何况，当多个候选节点保留了节点并准备绑定时) ，然后使用标准的Kubernetes函数/插件(例如，扩展与包首选项)对虚拟节点进行计分，最终只有一个虚拟节点被保留给代理节点。然后，代理调度器用`multicluster.admiralty.io/is-allowed=true`注释目标集群中的pod伴侣，通知候选调度器候选pod是委托pod。代理调度程序再次等待，这次是在代理pod的绑定周期开始时，查看是否成功绑定了委托pod。

接下来，候选调度程序绑定委托pod。无论成功还是失败，代理调度程序都可以从pod伴侣的状态看到它。如果成功，则删除所有其他候选pod。如果失败，则将其放回调度队列中。

### 总结

Pod伴侣注释被用作代理调度器和候选调度器之间的双向跨集群通讯通道，以协调调度和绑定周期。当调度一个代理Pod（将其绑定到一个虚拟节点）时，代理调度器对目标集群知之甚少。它不基于聚合数据进行过滤，因为这可能不准确。相反，它将候选Pod发送至所有目标集群。候选调度器拥有确定这些Pod是否可以被调度所需的所有情况。在基于聚合且足够的数据过滤通过的虚拟节点得分之后，代理调度器选出一个候选Pod作为委托Pod。最终，委托Pod被绑定，代理Pod也被绑定，而所有其他的候选Pod则被删除。

## 源码剖析

admiralty总共包含四个组件，由此挨个进行分析。

> 1.admiralty-multicluster-scheduler-candidate-scheduler
>
> 2.admiralty-multicluster-scheduler-controller-manager
>
> 3.admiralty-multicluster-scheduler-proxy-scheduler
>
> 4.admiralty-multicluster-scheduler-restarter

### admiralty-multicluster-scheduler-restarter

这个组件的用途是持续监听集群内的Scret和自定义资源Target，将每个时刻的Target与Secret进行匹配，并记录在当前组件上下文中。当上一时刻的`Target-Secret`对与当前不同，则重启代理调度器`proxy-scheduler`与控制管理器`proxy-scheduler`，由它们更新`Virtual nodes`列表并完成跨集群调度工作。

```go
func main() {
	stopCh := signals.SetupSignalHandler()
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<-stopCh
		cancel()
	}()

	cfg, ns, err := config.ConfigAndNamespaceForKubeconfigAndContext("", "")
	utilruntime.Must(err)

	k, err := kubernetes.NewForConfig(cfg) //获取kubeconfig
	utilruntime.Must(err)

	customClient, err := client.NewForConfig(cfg) //创建出一个与集群资源交互的客户端
	utilruntime.Must(err)

	kubeInformerFactory := kubeinformers.NewSharedInformerFactory(k, time.Second*30)          //创建 Kubernetes 的系统资源的 Informer 工厂
	customInformerFactory := informers.NewSharedInformerFactory(customClient, time.Second*30) //创建用户自定义的资源的 Informer 工厂
	//创建 Kubernetes 的 Informer 工厂，用于生产 Informer 对象。Informer 对象主要用于监听（watch）Kubernetes 中的资源变化。
	//这种机制可以高效地获取到集群中资源的实时状态信息，从而让我们能基于这些信息进行一系列的操作，如控制循环、事件触发等。

	targetCtrl := target.NewController(k, ns,
		os.Getenv("ADMIRALTY_CONTROLLER_MANAGER_DEPLOYMENT_NAME"),
		os.Getenv("ADMIRALTY_PROXY_SCHEDULER_DEPLOYMENT_NAME"),
		customInformerFactory.Multicluster().V1alpha1().ClusterTargets(), //创建一个用于监听 ClusterTargets 自定义资源的 Informer
		customInformerFactory.Multicluster().V1alpha1().Targets(),        //创建一个用于监听 Targets 自定义资源的 Informer
		kubeInformerFactory.Core().V1().Secrets())                        //创建一个用于监听核心 Kubernetes API 中 Secrets 资源的 Informer

	kubeInformerFactory.Start(stopCh)
	customInformerFactory.Start(stopCh)

	var leaderElect bool
	flag.BoolVar(&leaderElect, "leader-elect", false, "Start a leader election client and gain leadership before executing the main loop. Enable this when running replicated components for high availability.")
	flag.Parse()

	if leaderElect {
		leaderelection.Run(ctx, ns, "admiralty-restarter", k, func(ctx context.Context) {
			utilruntime.Must(targetCtrl.Run(1, stopCh))
		})
	} else {
		utilruntime.Must(targetCtrl.Run(1, stopCh)) //持续监控Target资源并与当前上下文所记录的进行对比，如有更新则重启其他调度器
	}
}
```

### admiralty-multicluster-scheduler-controller-manager

这个组件主要负责Virtual Node的生命周期管理，针对变异pod的WebHook服务创建，创建服务使得源集群可以对pod进行log,exec等操作。

```go
func main() {
	stopCh := signals.SetupSignalHandler()
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<-stopCh
		cancel()
	}()
	o := parseFlags()
	setupLogging(ctx, o)
	agentCfg := agentconfig.NewFromCRD(ctx)                                  //agentCfg 包含所有Target信息的KubeConfig
	cfg, ns, err := config.ConfigAndNamespaceForKubeconfigAndContext("", "") //KubeConfig
	utilruntime.Must(err)
	k, err := kubernetes.NewForConfig(cfg) //KubeConfig Client
	utilruntime.Must(err)
	startWebhook(stopCh, cfg)                       //启动一个 Kubernetes Admission Webhook 服务，允许你在 Kubernetes API Server 处理请求之前或之后拦截这些请求，通过自定义逻辑修改或拒绝资源
	go startVirtualKubeletServers(ctx, agentCfg, k) //创建一个 Server 用于处理针对 pod 的 log,exec 请求

	if o.leaderElect {
		leaderelection.Run(ctx, ns, "admiralty-controller-manager", k, func(ctx context.Context) {
			runControllers(ctx, stopCh, agentCfg, cfg, k)
		})
	} else {
		runControllers(ctx, stopCh, agentCfg, cfg, k) //将Targets数据转化为Virtual Node
	}
}
```

### admiralty-multicluster-scheduler-proxy-scheduler & admiralty-multicluster-scheduler-candidate-scheduler

根据镜像运行参数选择某一个插件运行调度程序，它们各自实现了部分`k8s.io/kubernetes/pkg/scheduler/framework`调度组件中的接口。

```go
package main

import (
	"math/rand"
	"os"
	"time"

	"admiralty.io/multicluster-scheduler/pkg/scheduler_plugins/candidate"
	"admiralty.io/multicluster-scheduler/pkg/scheduler_plugins/proxy"
	"github.com/spf13/pflag"
	cliflag "k8s.io/component-base/cli/flag"
	"k8s.io/component-base/logs"
	scheduler "k8s.io/kubernetes/cmd/kube-scheduler/app"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// BEWARE candidate and proxy must run in different processes, because a scheduler only processes one pod at a time
	// and proxy waits on candidates in filter plugin

	command := scheduler.NewSchedulerCommand(
		scheduler.WithPlugin(candidate.Name, candidate.New),
		scheduler.WithPlugin(proxy.Name, proxy.New))

	// TODO: once we switch everything over to Cobra commands, we can go back to calling
	// utilflag.InitFlags() (by removing its pflag.Parse() call). For now, we have to set the
	// normalize func and add the go flag set by hand.
	pflag.CommandLine.SetNormalizeFunc(cliflag.WordSepNormalizeFunc)
	// utilflag.InitFlags()
	logs.InitLogs()
	defer logs.FlushLogs()

	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
```

```go
// proxy/plugin.go
type Plugin struct {
	handle  framework.FrameworkHandle
	targets map[string]*versioned.Clientset

	failedNodeNamesByPodUID map[types.UID]map[string]bool
	mx                      sync.RWMutex
}

var _ framework.FilterPlugin = &Plugin{}
var _ framework.ReservePlugin = &Plugin{}
var _ framework.UnreservePlugin = &Plugin{}
var _ framework.PreBindPlugin = &Plugin{}
var _ framework.PostBindPlugin = &Plugin{}
```

```go
// candidate/plugin.go
type Plugin struct {
	handle framework.FrameworkHandle
	client versioned.Interface
}

var _ framework.PreFilterPlugin = &Plugin{}
var _ framework.ReservePlugin = &Plugin{}
var _ framework.PreBindPlugin = &Plugin{}
```