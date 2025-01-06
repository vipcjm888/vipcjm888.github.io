# 云原生应用之路——从 Kubernetes 到云原生

注：本文根据笔者在 [ArchSummit 2017 北京站](https://archsummit.infoq.cn/2017/beijing/presentation/306) 和数人云 & TalkingData 合办的 Service Mesh is coming meetup 中分享的话题《从 Kubernetes 到云原生—— 云原生应用之路》改写而成。

本文简要介绍了容器技术发展的路径，为何 Kubernetes 的出现是容器技术发展到这一步的必然选择，而为何 Kubernetes 又将成为云原生应用的基石。

我的分享按照这样的主线展开：容器 -> Kubernetes -> 微服务 ->云原生 -> 服务网格 -> 使用场景 -> 开源。

## 容器

> 容器 ——云原生的基石

容器最初是通过开发者工具而流行，可以使用它来做隔离的开发测试环境和持续集成环境，这些都是因为容器轻量级，易于配置和使用带来的优势，docker 和 docker-compose 这样的工具极大的方便的了应用开发环境的搭建，开发者就像是化学家一样在其中小心翼翼的进行各种调试和开发。

随着容器的在开发者中的普及，已经大家对 CI 流程的熟悉，容器周边的各种工具蓬勃发展，俨然形成了一个小生态，在 2016 年达到顶峰，下面这张是我画的容器生态图。

![容器生态图](../images/container-ecosystem.png)

该生态涵盖了容器应用中从镜像仓库、服务编排、安全管理、持续集成与发布、存储和网络管理等各个方面，随着在单主机中运行容器的成熟，集群管理和容器编排成为容器技术亟待解决的问题。譬如化学家在实验室中研究出来的新产品，如何推向市场，进行大规模生产，成了新的议题。

## 为什么使用 Kubernetes

> Kubernetes—— 让容器应用进入大规模工业生产。

**Kubernetes 是容器编排系统的事实标准**

在单机上运行容器，无法发挥它的最大效能，只有形成集群，才能最大程度发挥容器的良好隔离、资源分配与编排管理的优势，而对于容器的编排管理，Swarm、Mesos 和 Kubernetes 的大战已经基本宣告结束，Kubernetes 成为了无可争议的赢家。

下面这张图是 Kubernetes 的架构图（图片来自网络），其中显示了组件之间交互的接口 CNI、CRI、OCI 等，这些将 Kubernetes 与某款具体产品解耦，给用户最大的定制程度，使得 Kubernetes 有机会成为跨云的真正的云原生应用的操作系统。

![Kubernetes架构](../images/kubernetes-high-level-component-archtecture.jpg)

随着 Kubernetes 的日趋成熟，“Kubernetes is becoming boring”，基于该 “操作系统” 之上构建的适用于不同场景的应用将成为新的发展方向，就像我们将石油开采出来后，提炼出汽油、柴油、沥青等等，所有的材料都将找到自己的用途，Kubernetes 也是，毕竟我们谁也不是为了部署和管理容器而用 Kubernetes，承载其上的应用才是价值之所在。

**云原生的核心目标**

![云原生的核心目标](../images/cloud-native-core-target.jpg)

云已经可以为我们提供稳定可以唾手可得的基础设施，但是业务上云成了一个难题，Kubernetes 的出现与其说是从最初的容器编排解决方案，倒不如说是为了解决应用上云（即云原生应用）这个难题。

包括微服务和 FaaS/Serverless 架构，都可以作为云原生应用的架构。

![FaaS Landscape](../images/redpoint-faas-landscape.jpg)

但就 2017 年为止，Kubernetes 的主要使用场景也主要作为应用开发测试环境、CI/CD 和运行 Web 应用这几个领域，如下图 [TheNewStack](http://thenewstack.io/) 的 Kubernetes 生态状况调查报告所示。

![运行在 Kubernetes 上的负载（2017 年）](../images/0069RVTdgy1fv5mxr6fxtj31kw11q484.jpg)

另外基于 Kubernetes 的构建 PaaS 平台和 Serverless 也处于爆发的准备的阶段，如下图中 Gartner 的报告中所示：

![Gartner 技术爆发趋势图（2017 年）](../images/0069RVTdgy1fv5my2jtxzj315o0z8dkr.jpg)

2017 年时各大公有云如 Google GKE、微软 Azure ACS、亚马逊 EKS（2018 年上线）、VMware、Pivotal（后被VMware 收购）、腾讯云、阿里云等都提供了 Kubernetes 服务。

## 微服务

> 微服务——Cloud Native 的应用架构。

下图是 [Bilgin Ibryam](https://developers.redhat.com/blog/author/bibryam/) 给出的微服务中应该关心的主题，图片来自 [RedHat Developers](https://developers.redhat.com/blog/2016/12/09/spring-cloud-for-microservices-compared-to-kubernetes/)。

![微服务的关注点](../images/microservices-concerns.jpg)

微服务带给我们很多开发和部署上的灵活性和技术多样性，但是也增加了服务调用的开销、分布式系统管理、调试与服务治理方面的难题。

当前最成熟最完整的微服务框架可以说非 [Spring](https://spring.io/) 莫属，而 Spring 又仅限于 Java 语言开发，其架构本身又跟 Kubernetes 存在很多重合的部分，如何探索将 Kubernetes 作为微服务架构平台就成为一个热点话题。

就拿微服务中最基础的**服务注册发现**功能来说，其方式分为**客户端服务发现**和**服务端服务发现**两种，Java 应用中常用的方式是使用 Eureka 和 Ribbon 做服务注册发现和负载均衡，这属于客户端服务发现，而在 Kubernetes 中则可以使用 DNS、Service 和 Ingress 来实现，不需要修改应用代码，直接从网络层面来实现。

![微服务中的两种服务发现方式](../images/service-discovery-in-microservices.png)

## 云原生

> DevOps——通向云原生的云梯

CNCF（云原生计算基金会）给出了云原生应用的三大特征：

- **容器化包装**：软件应用的进程应该包装在容器中独立运行。
- **动态管理**：通过集中式的编排调度系统来动态的管理和调度。
- **微服务化**：明确服务间的依赖，互相解耦。

下图是我整理的关于云原生所需要的能力和特征。

![Cloud Native Features](https://jimmysong.io/kubernetes-handbook/images/cloud-native-architecutre-mindnode.jpg)

[CNCF](https://cncf.io) 所托管的应用（即正式捐献给 CNCF 的应用，2017 年已达 12 个），即朝着这个目标发展，其公布的 [Cloud Native Landscape](https://github.com/cncf/landscape)，给出了云原生生态的参考体系。

![Cloud Native Landscape v1.0](../images/0069RVTdgy1fv5myp6ednj31kw0w0u0x.jpg)

**使用 Kubernetes 构建云原生应用**

我们都是知道 Heroku 推出了适用于 PaaS 的 [12 因素应用](https://12factor.net/)的规范，包括如下要素：

1. 基准代码
2. 依赖管理
3. 配置
4. 后端服务
5. 构建，发布，运行
6. 无状态进程
7. 端口绑定
8. 并发
9. 易处理
10. 开发环境与线上环境等价
11. 日志作为事件流
12. 管理进程

另外还有补充的三点：

- API声明管理
- 认证和授权
- 监控与告警

如果落实的具体的工具，请看下图，使用Kubernetes构建云原生架构：

![使用 Kubernetes 构建云原生架构应用](../images/building-cloud-native-architecture-with-kubernetes.png)

结合这 12 因素对开发或者改造后的应用适合部署到 Kubernetes 之上，基本流程如下图所示：

![创建 Kubernetes 原生应用](../images/creating-kubernetes-native-app.jpg)

**迁移到云架构**

迁移到云端架构，相对单体架构来说会带来很多挑战。比如自动的持续集成与发布、服务监控的变革、服务暴露、权限的管控等。这些具体细节请参考 [Kubernetes Handbook](https://jimmysong.io/kubernetes-handbook) 中的说明，在此就不细节展开，另外推荐一本我翻译的由 Pivotal 出品的电子书——《迁移到云原生应用架构》，推荐大家阅读。

## 服务网格

> Services for show, meshes for a pro.

Kubernetes 中的应用将作为微服务运行，但是 Kubernetes 本身并没有给出微服务治理的解决方案，比如服务的限流、熔断、良好的灰度发布支持等。

**Service Mesh 可以用来做什么**

- Traffic Management：API 网关
- Observability：服务调用和性能分析
- Policy Enforcment：控制服务访问策略
- Service Identity and Security：安全保护

**Service Mesh 的特点**

- 专用的基础设施层
- 轻量级高性能网络代理
- 提供安全的、快速的、可靠地服务间通讯
- 扩展 kubernetes 的应用负载均衡机制，实现灰度发布
- 完全解耦于应用，应用可以无感知，加速应用的微服务和云原生转型

使用 Service Mesh 将可以有效的治理 Kubernetes 中运行的服务，当前开源的流行的 Service Mesh 有：

- [Linkerd](https://linkerd.io/)：由最早提出 Service Mesh 的公司 [Buoyant](https://buoyant.io/) 开源，创始人来自 Twitter
- [Envoy](https://www.envoyproxy.io/)：由 Lyft 开源，可以在 Istio 中使用 Sidecar 模式运行以作为数据平面，也可以基于它来构建自己的服务网格
- [Istio](https://istio.io)：由 Google、IBM、Lyft 联合开发并开源

此外还有很多其它的 Service Mesh 鱼贯而出，请参考 [awesome-cloud-native](https://jimmysong.io/awesome-cloud-native)。

**Istio VS Linkerd**

Linkerd 和 Istio 是最早开源的 Service Mesh，它们都支持 Kubernetes，下面是它们之间的一些特性对比。

| **Feature**      | **Istio**     | **Linkerd**                  |
| ---------------- | ------------- | ---------------------------- |
| 部署架构         | Envoy/Sidecar | DaemonSets                   |
| 易用性           | 复杂          | 简单                         |
| 支持平台         | Kubernetes    | Kubernetes/Mesos/Istio/Local |
| 是否已有生产部署 | 是            | 是                           |

关于两者的架构可以参考各自的官方文档，我只从其在Kubernetes上的部署结构来说明其区别。

![istio vs linkerd](../images/istio-vs-linkerd.jpg)

Istio 的组件复杂，可以分别部署在 Kubernetes 集群中，但是作为核心路由组件 **Envoy** 是以 **Sidecar** 形式与应用运行在同一个 Pod 中的，所有进入该 Pod 中的流量都需要先经过 Envoy。

Linker 的部署十分简单，本身就是一个镜像，使用 Kubernetes 的 [DaemonSet](https://jimmysong.io/kubernetes-handbook/concepts/daemonset.html) 方式在每个 node 节点上运行。

## 使用场景

> 云原生的大规模工业生产

**GitOps**

给开发者带来最大的配置和上线的灵活性，践行 DevOps 流程，改善研发效率，下图这样的情况将更少发生。

![部署流水线](../images/0069RVTdgy1fv5mzj8rj6j318g1ewtfc.jpg)

我们知道 Kubernetes 中的所有应用的部署都是基于YAML文件的，这实际上就是一种 **Infrastructure as code**，完全可以通过 Git 来管控基础设施和部署环境的变更。

**大数据**

Spark 现在已经非官方支持了基于 Kubernetes 的原生调度，其具有以下特点：

- Kubernetes 原生调度：与 yarn、mesos 同级
- 资源隔离，粒度更细：以 namespace 来划分用户
- 监控的变革：单次任务资源计量
- 日志的变革：pod 的日志收集

| **特性** | **Yarn**         | **Kubernetes** |
| -------- | ---------------- | -------------- |
| 队列     | queue            | namespace      |
| 实例     | ExcutorContainer | Executor Pod   |
| 网络     | host             | plugin         |
| 异构     | no               | yes            |
| 安全     | RBAC             | ACL            |

下图是在 Kubernetes 上运行三种调度方式的 spark 的单个节点的应用部分对比：

![使用不同调度器的 Spark on Kubernetes](../images/spark-on-kubernetes-with-different-schedulers.jpg)

从上图中可以看到在 Kubernetes 上使用 YARN 调度、standalone 调度和 Kubernetes 原生调度的方式，每个 node 节点上的 Pod 内的 Spark Executor 分布，毫无疑问，使用 Kubernetes 原生调度的 Spark 任务才是最节省资源的。

提交任务的语句看起来会像是这样的：

```bash
./spark-submit \
  --deploy-mode cluster \
  --class com.talkingdata.alluxio.hadooptest \
  --master k8s://https://172.20.0.113:6443 \
  --kubernetes-namespace spark-cluster \
  --conf spark.kubernetes.driverEnv.SPARK_USER=hadoop \
  --conf spark.kubernetes.driverEnv.HADOOP_USER_NAME=hadoop \
  --conf spark.executorEnv.HADOOP_USER_NAME=hadoop \
  --conf spark.executorEnv.SPARK_USER=hadoop \
  --conf spark.kubernetes.authenticate.driver.serviceAccountName=spark \
  --conf spark.driver.memory=100G \
  --conf spark.executor.memory=10G \
  --conf spark.driver.cores=30 \
  --conf spark.executor.cores=2 \
  --conf spark.driver.maxResultSize=10240m \
  --conf spark.kubernetes.driver.limit.cores=32 \
  --conf spark.kubernetes.executor.limit.cores=3 \
  --conf spark.kubernetes.executor.memoryOverhead=2g \
  --conf spark.executor.instances=5 \
  --conf spark.app.name=spark-pi \
  --conf spark.kubernetes.driver.docker.image=spark-driver:v2.1.0-kubernetes-0.3.1-1 \
  --conf spark.kubernetes.executor.docker.image=spark-executor:v2.1.0-kubernetes-0.3.1-1 \
  --conf spark.kubernetes.initcontainer.docker.image=spark-init:v2.1.0-kubernetes-0.3.1-1 \
  --conf spark.kubernetes.resourceStagingServer.uri=http://172.20.0.114:31000 \
~/Downloads/tendcloud_2.10-1.0.jar
```

关于支持 Kubernetes 原生调度的 Spark 请参考 [spark-on-k8s](https://jimmysong.io/spark-on-k8s/)。

## 开源

> Contributing is Not only about code, it is about helping a community.

下图是我们刚调研准备使用 Kubernetes 时候的调研方案选择。

![Kubernetes 解决方案](../images/0069RVTdgy1fv5mzywc83j31fk1i8qg4.jpg)

对于一个初次接触 Kubernetes 的人来说，看到这样一个庞大的架构选型时会望而生畏，但是 Kubernetes 的开源社区帮助了我们很多。

![Kubernetes SIG](../images/kubernetes-sigs.jpg)

## 更多

Bilgin Ibryam 写了这篇《[分布式系统在 Kubernetes 上的进化](https://cloudnative.to/blog/distributed-systems-kubernetes/)》，可以帮助大家更好的理解基于 Kubernetes 的分布式系统的演进。