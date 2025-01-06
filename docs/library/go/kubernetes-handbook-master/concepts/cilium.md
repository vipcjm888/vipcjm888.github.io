# 基于 eBPF 的网络 Cilium

Cilium 是一款开源软件，也是 CNCF 的孵化项目，目前[已有公司](https://isovalent.com/)提供商业化支持，还有基于 Cilium 实现的服务网格解决方案。最初它仅是作为一个 Kubernetes 网络组件。Cilium 在 1.7 版本后[推出并开源了 Hubble](https://cilium.io/blog/2019/11/19/announcing-hubble)，它是专门为网络可视化设计，能够利用 Cilium 提供的 eBPF 数据路径，获得对 Kubernetes 应用和服务的网络流量的深度可见性。这些网络流量信息可以对接 Hubble CLI、UI 工具，可以通过交互式的方式快速进行问题诊断。除了 Hubble 自身的监控工具，还可以对接主流的云原生监控体系——Prometheus 和 Grafana，实现可扩展的监控策略。

![Cilium](../images/006tNbRwly1fwqi98i51ij30sc0j80zn.jpg)

本节将带你了解什么是 Cilium 及选择它的原因。

## Cilium 是什么？

Cilium 为基于 Kubernetes 的 Linux 容器管理平台上部署的服务，透明地提供服务间的网络和 API 连接及安全。

Cilium 底层是基于 Linux 内核的新技术 eBPF，可以在 Linux 系统中动态注入强大的安全性、可见性和网络控制逻辑。 Cilium 基于 eBPF 提供了多集群路由、替代 kube-proxy 实现负载均衡、透明加密以及网络和服务安全等诸多功能。除了提供传统的网络安全之外，eBPF 的灵活性还支持应用协议和 DNS 请求/响应安全。同时，Cilium 与 Envoy 紧密集成，提供了基于 Go 的扩展框架。因为 eBPF 运行在 Linux 内核中，所以应用所有 Cilium 功能，无需对应用程序代码或容器配置进行任何更改。

基于微服务的应用程序分为小型独立服务，这些服务使用 **HTTP**、**gRPC**、**Kafka** 等轻量级协议通过 API 相互通信。但是，现有的 Linux 网络安全机制（例如 iptables）仅在网络和传输层（即 IP 地址和端口）上运行，并且缺乏对微服务层的可见性。

Cilium 为 Linux 容器框架（如 [**Docker**](https://www.docker.com/) 和 [**Kubernetes）**](https://kubernetes.io/) 带来了 API 感知网络安全过滤。使用名为 **eBPF** 的新 Linux 内核技术，Cilium 提供了一种基于容器 / 容器标识定义和实施网络层和应用层安全策略的简单而有效的方法。

**注**：Cilium 中文意思是 “纤毛 “，它十分细小而又无处不在。

> ## eBPF
>
> **扩展的柏克莱封包过滤器**（extented Berkeley Packet Filter，缩写 eBPF），是 [类 Unix](https://zh.wikipedia.org/wiki/%E7%B1%BBUnix) 系统上 [数据链路层](https://zh.wikipedia.org/wiki/%E6%95%B0%E6%8D%AE%E9%93%BE%E8%B7%AF%E5%B1%82) 的一种原始接口，提供原始链路层 [封包](https://zh.wikipedia.org/wiki/%E5%B0%81%E5%8C%85) 的收发，除此之外，如果网卡驱动支持 [洪泛](https://zh.wikipedia.org/wiki/%E6%B4%AA%E6%B3%9B) 模式，那么它可以让网卡处于此种模式，这样可以收到 [网络](https://zh.wikipedia.org/wiki/%E7%BD%91%E7%BB%9C) 上的所有包，不管他们的目的地是不是所在 [主机](https://zh.wikipedia.org/wiki/%E4%B8%BB%E6%A9%9F)。参考 [维基百科](https://zh.wikipedia.org/wiki/BPF) 和 [BPF、eBPF、XDP 和 Bpfilter 的区别](https://www.netronome.com/blog/bpf-ebpf-xdp-and-bpfilter-what-are-these-things-and-what-do-they-mean-enterprise/)。

## Hubble 是什么？

Hubble 是一个完全分布式的网络和安全可观察性平台。它建立在 Cilium 和 eBPF 之上，以完全透明的方式实现对服务的通信和行为以及网络基础设施的深度可见性（visibility）。

通过建立在 Cilium 之上，Hubble 可以利用 eBPF 实现可见性。依靠 eBPF，所有的可见性都是可编程的，并允许采用一种动态方法，最大限度地减少开销，同时按照用户的要求提供深入和详细的可见性。Hubble 的创建和专门设计是为了最好地利用 eBPF 的能力。

## 特性

以下是 Cilium 的特性。

**基于身份的安全性**

Cilium 可见性和安全策略基于容器编排系统的标识（例如，Kubernetes 中的 Label）。在编写安全策略、审计和故障排查时，再也不用担心网络子网或容器 IP 地址了。

**卓越的性能**

eBPF 利用 Linux 底层的强大能力，通过提供 Linux 内核的沙盒可编程性来实现数据路径，从而提供卓越的性能。

**API 协议可见性 + 安全性 **

传统防火墙仅根据 IP 地址和端口等网络标头查看和过滤数据包。Cilium 也可以这样做，但也可以理解并过滤单个 HTTP、gRPC 和 Kafka 请求，这些请求将微服务拼接在一起。

**专为扩展而设计**

Cilium 是为扩展而设计的，在部署新 pod 时不需要节点间交互，并且通过高度可扩展的键值存储进行所有协调。

## 为什么选择 Cilium 和 Hubble？

现代数据中心应用程序的开发已经转向面向服务的体系结构（SOA），通常称为微服务，其中大型应用程序被分成小型独立服务，这些服务使用 HTTP 等轻量级协议通过 API 相互通信。微服务应用程序往往是高度动态的，作为持续交付的一部分部署的滚动更新期间单个容器启动或销毁，应用程序扩展 / 缩小以适应负载变化。

这种向高度动态的微服务的转变过程，给确保微服务之间的连接方面提出了挑战和机遇。传统的 Linux 网络安全方法（例如 iptables）过滤 IP 地址和 TCP/UDP 端口，但 IP 地址经常在动态微服务环境中流失。容器的高度不稳定的生命周期导致这些方法难以与应用程序并排扩展，因为负载均衡表和访问控制列表要不断更新，可能增长成包含数十万条规则。出于安全目的，协议端口（例如，用于 HTTP 流量的 TCP 端口 80）不能再用于区分应用流量，因为该端口用于跨服务的各种消息。

另一个挑战是提供准确的可见性，因为传统系统使用 IP 地址作为主要识别工具，其在微服务架构中的寿命可能才仅仅几秒钟，被大大缩短。

利用 Linux eBPF，Cilium 保留了透明地插入安全可视性 + 强制执行的能力，但这种方式基于服务 /pod/ 容器标识（与传统系统中的 IP 地址识别相反），并且可以根据应用层进行过滤 （例如 HTTP）。因此，通过将安全性与寻址分离，Cilium 不仅可以在高度动态的环境中应用安全策略，而且除了提供传统的第 3 层和第 4 层分割之外，还可以通过在 HTTP 层运行来提供更强的安全隔离。 

eBPF 的使用使得 Cilium 能够以高度可扩展的方式实现以上功能，即使对于大规模环境也不例外。

## 功能概述

### 透明的保护 API

能够保护现代应用程序协议，如 REST/HTTP、gRPC 和 Kafka。传统防火墙在第 3 层和第 4 层运行，在特定端口上运行的协议要么完全受信任，要么完全被阻止。Cilium 提供了过滤各个应用程序协议请求的功能，例如：

- 允许所有带有方法 `GET` 和路径 `/public/.*` 的 HTTP 请求。拒绝所有其他请求。
- 允许 `service1` 在 Kafka topic 上生成 `topic1`，`service2` 消费 `topic1`。拒绝所有其他 Kafka 消息。
- 要求 HTTP 标头 `X-Token: [0-9]+` 出现在所有 REST 调用中。

详情请参考 [7 层协议](http://docs.cilium.io/en/stable/policy/#layer-7)。

### 基于身份来保护服务间通信

现代分布式应用程序依赖于诸如容器之类的技术来促进敏捷性并按需扩展。这将导致在短时间内启动大量应用容器。典型的容器防火墙通过过滤源 IP 地址和目标端口来保护工作负载。这就要求不论在集群中的哪个位置启动容器时都要操作所有服务器上的防火墙。

为了避免受到规模限制，Cilium 为共享相同安全策略的应用程序容器组分配安全标识。然后，该标识与应用程序容器发出的所有网络数据包相关联，从而允许验证接收节点处的身份。使用键值存储执行安全身份管理。

### 安全访问外部服务

基于标签的安全性是集群内部访问控制的首选工具。为了保护对外部服务的访问，支持入口（ingress）和出口（egress）的传统基于 CIDR 的安全策略。这允许限制对应用程序容器的访问以及对特定 IP 范围的访问。

### 简单网络

一个简单的扁平第 3 层网络能够跨越多个集群连接所有应用程序容器。使用主机范围分配器可以简化 IP 分配。这意味着每个主机可以在主机之间没有任何协调的情况下分配 IP。

支持以下多节点网络模型：

- **Overlay**：基于封装的虚拟网络产生所有主机。目前 VXLAN 和 Geneve 已经完成，但可以启用 Linux 支持的所有封装格式。

  何时使用此模式：此模式具有最小的基础架构和集成要求。它几乎适用于任何网络基础架构，唯一的要求是主机之间可以通过 IP 连接。

- **本机路由**：使用 Linux 主机的常规路由表。网络必须能够路由应用程序容器的 IP 地址。

  何时使用此模式：此模式适用于高级用户，需要了解底层网络基础结构。此模式适用于：

  - 本地 IPv6 网络
  - 与云网络路由器配合使用
  - 如果您已经在运行路由守护进程

### 负载均衡

应用程序容器和外部服务之间的流量的分布式负载均衡。负载均衡使用 eBPF 实现，允许几乎无限的规模，并且如果未在源主机上执行负载均衡操作，则支持直接服务器返回（DSR）。 

**注意**：负载均衡需要启用连接跟踪。这是默认值。

### 监控和故障排除

可见性和故障排查是任何分布式系统运行的基础。虽然我们喜欢用 `tcpdump` 和 `ping`，它们很好用，但我们努力为故障排除提供更好的工具。包括以下工具：

- 使用元数据进行事件监控：当数据包被丢弃时，该工具不仅仅报告数据包的源 IP 和目标 IP，该工具还提供发送方和接收方的完整标签信息等。
- 策略决策跟踪：为什么丢弃数据包或拒绝请求。策略跟踪框架允许跟踪运行工作负载和基于任意标签定义的策略决策过程。
- 通过 Prometheus 导出指标：通过 Prometheus 导出关键指标，以便与现有仪表板集成。

### 集成

- 网络插件集成：[CNI](https://github.com/containernetworking/cni)、[libnetwork](https://github.com/docker/libnetwork)
- 容器运行时：[containerd](https://github.com/containerd/containerd)
- Kubernetes：[NetworkPolicy](https://kubernetes.io/docs/concepts/services-networking/network-policies/)、[Label](https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/)、[Ingress](https://kubernetes.io/docs/concepts/services-networking/ingress/)、[Service](https://kubernetes.io/docs/concepts/services-networking/service/)
- 日志记录：syslog、[fluentd](http://www.fluentd.org/)

## 参考

- [Cilium 官方网站 - cilium.io](https://cilium.io)
- [eBPF 简史 - ibm.com](https://www.ibm.com/developerworks/cn/linux/l-lo-eBPF-history/index.html)
- [网络层拦截可选项 - zhihu.com](https://zhuanlan.zhihu.com/p/25672552)
- [BPF, eBPF, XDP and Bpfilter… What are These Things and What do They Mean for the Enterprise? - netronome.com](https://www.netronome.com/blog/bpf-ebpf-xdp-and-bpfilter-what-are-these-things-and-what-do-they-mean-enterprise/)
