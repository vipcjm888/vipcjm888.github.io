**Rust 语言圣经**

- [关于本书](/markdown/rust/rust-course/index.md)
- [进入 Rust 编程世界](/markdown/rust/rust-course/into-rust.md)
- [避免从入门到放弃](/markdown/rust/rust-course/first-try/sth-you-should-not-do.md)
- [快速查询入口](/markdown/rust/rust-course/index-list.md)

---
- [Xobserve: 一切皆可观测](/markdown/rust/rust-course/some-thoughts.md)
- [Rust 翻译计划( 代号 Rustt )](/markdown/rust/rust-course/rustt.md)

**Rust 语言基础学习**

---

- [寻找牛刀，以便小试](/markdown/rust/rust-course/first-try/intro.md)
  - [安装 Rust 环境](/markdown/rust/rust-course/first-try/installation.md)
  - [墙推 VSCode!](/markdown/rust/rust-course/first-try/editor.md)
  - [认识 Cargo](/markdown/rust/rust-course/first-try/cargo.md)
  - [不仅仅是 Hello world](/markdown/rust/rust-course/first-try/hello-world.md)
  - [下载依赖太慢了？](/markdown/rust/rust-course/first-try/slowly-downloading.md)

- [Rust 基础入门](/markdown/rust/rust-course/basic/intro.md)

  - [变量绑定与解构](/markdown/rust/rust-course/basic/variable.md)
  - [基本类型](/markdown/rust/rust-course/basic/base-type/index.md)
    - [数值类型](/markdown/rust/rust-course/basic/base-type/numbers.md)
    - [字符、布尔、单元类型](/markdown/rust/rust-course/basic/base-type/char-bool.md)
    - [语句与表达式](/markdown/rust/rust-course/basic/base-type/statement-expression.md)
    - [函数](/markdown/rust/rust-course/basic/base-type/function.md)
  - [所有权和借用](/markdown/rust/rust-course/basic/ownership/index.md)
    - [所有权](/markdown/rust/rust-course/basic/ownership/ownership.md)
    - [引用与借用](/markdown/rust/rust-course/basic/ownership/borrowing.md)
  - [复合类型](/markdown/rust/rust-course/basic/compound-type/intro.md)
    - [字符串与切片](/markdown/rust/rust-course/basic/compound-type/string-slice.md)
    - [元组](/markdown/rust/rust-course/basic/compound-type/tuple.md)
    - [结构体](/markdown/rust/rust-course/basic/compound-type/struct.md)
    - [枚举](/markdown/rust/rust-course/basic/compound-type/enum.md)
    - [数组](/markdown/rust/rust-course/basic/compound-type/array.md)
  - [流程控制](/markdown/rust/rust-course/basic/flow-control.md)
  - [模式匹配](/markdown/rust/rust-course/basic/match-pattern/intro.md)
    - [match 和 if let](/markdown/rust/rust-course/basic/match-pattern/match-if-let.md)
    - [解构 Option](/markdown/rust/rust-course/basic/match-pattern/option.md)
    - [模式适用场景](/markdown/rust/rust-course/basic/match-pattern/pattern-match.md)
    - [全模式列表](/markdown/rust/rust-course/basic/match-pattern/all-patterns.md)
  - [方法 Method](/markdown/rust/rust-course/basic/method.md)
  - [泛型和特征](/markdown/rust/rust-course/basic/trait/intro.md)
    - [泛型 Generics](/markdown/rust/rust-course/basic/trait/generic.md)
    - [特征 Trait](/markdown/rust/rust-course/basic/trait/trait.md)
    - [特征对象](/markdown/rust/rust-course/basic/trait/trait-object.md)
    - [进一步深入特征](/markdown/rust/rust-course/basic/trait/advance-trait.md)
  - [集合类型](/markdown/rust/rust-course/basic/collections/intro.md)
    - [动态数组 Vector](/markdown/rust/rust-course/basic/collections/vector.md)
    - [KV 存储 HashMap](/markdown/rust/rust-course/basic/collections/hashmap.md)
  - [认识生命周期](/markdown/rust/rust-course/basic/lifetime.md)
  - [返回值和错误处理](/markdown/rust/rust-course/basic/result-error/intro.md)
    - [panic! 深入剖析](/markdown/rust/rust-course/basic/result-error/panic.md)
    - [返回值 Result 和?](/markdown/rust/rust-course/basic/result-error/result.md)
  - [包和模块](/markdown/rust/rust-course/basic/crate-module/intro.md)
    - [包 Crate](/markdown/rust/rust-course/basic/crate-module/crate.md)
    - [模块 Module](/markdown/rust/rust-course/basic/crate-module/module.md)
    - [使用 use 引入模块及受限可见性](/markdown/rust/rust-course/basic/crate-module/use.md)
  - [注释和文档](/markdown/rust/rust-course/basic/comment.md)
  - [格式化输出](/markdown/rust/rust-course/basic/formatted-output.md)
- [入门实战：文件搜索工具](/markdown/rust/rust-course/basic-practice/intro.md)
  - [基本功能](/markdown/rust/rust-course/basic-practice/base-features.md)
  - [增加模块化和错误处理](/markdown/rust/rust-course/basic-practice/refactoring.md)
  - [测试驱动开发](/markdown/rust/rust-course/basic-practice/tests.md)
  - [使用环境变量](/markdown/rust/rust-course/basic-practice/envs.md)
  - [重定向错误信息的输出](/markdown/rust/rust-course/basic-practice/stderr.md)
  - [使用迭代器来改进程序(可选)](/markdown/rust/rust-course/basic-practice/iterators.md)

**Rust 语言进阶学习**

---

- [Rust 高级进阶](/markdown/rust/rust-course/advance/intro.md)
  - [生命周期](/markdown/rust/rust-course/advance/lifetime/intro.md)
    - [深入生命周期](/markdown/rust/rust-course/advance/lifetime/advance.md)
    - [&'static 和 T: 'static](/markdown/rust/rust-course/advance/lifetime/static.md)
    <!-- - [一些关于生命周期的误解 todo](/markdown/rust/rust-course/advance/lifetime/misconceptions.md) -->
  - [函数式编程: 闭包、迭代器](/markdown/rust/rust-course/advance/functional-programing/intro.md)
    - [闭包 Closure](/markdown/rust/rust-course/advance/functional-programing/closure.md)
    - [迭代器 Iterator](/markdown/rust/rust-course/advance/functional-programing/iterator.md)
  - [深入类型](/markdown/rust/rust-course/advance/into-types/intro.md)
    - [类型转换](/markdown/rust/rust-course/advance/into-types/converse.md)
    - [newtype 和 类型别名](/markdown/rust/rust-course/advance/into-types/custom-type.md)
    - [Sized 和不定长类型 DST](/markdown/rust/rust-course/advance/into-types/sized.md)
    - [枚举和整数](/markdown/rust/rust-course/advance/into-types/enum-int.md)
  - [智能指针](/markdown/rust/rust-course/advance/smart-pointer/intro.md)
    - [Box<T>堆对象分配](/markdown/rust/rust-course/advance/smart-pointer/box.md)
    - [Deref 解引用](/markdown/rust/rust-course/advance/smart-pointer/deref.md)
    - [Drop 释放资源](/markdown/rust/rust-course/advance/smart-pointer/drop.md)
    - [Rc 与 Arc 实现 1vN 所有权机制](/markdown/rust/rust-course/advance/smart-pointer/rc-arc.md)
    - [Cell 与 RefCell 内部可变性](/markdown/rust/rust-course/advance/smart-pointer/cell-refcell.md)
  - [循环引用与自引用](/markdown/rust/rust-course/advance/circle-self-ref/intro.md)
    - [Weak 与循环引用](/markdown/rust/rust-course/advance/circle-self-ref/circle-reference.md)
    - [结构体中的自引用](/markdown/rust/rust-course/advance/circle-self-ref/self-referential.md)
  - [多线程并发编程](/markdown/rust/rust-course/advance/concurrency-with-threads/intro.md)
    - [并发和并行](/markdown/rust/rust-course/advance/concurrency-with-threads/concurrency-parallelism.md)
    - [使用多线程](/markdown/rust/rust-course/advance/concurrency-with-threads/thread.md)
    - [线程同步：消息传递](/markdown/rust/rust-course/advance/concurrency-with-threads/message-passing.md)
    - [线程同步：锁、Condvar 和信号量](/markdown/rust/rust-course/advance/concurrency-with-threads/sync1.md)
    - [线程同步：Atomic 原子操作与内存顺序](/markdown/rust/rust-course/advance/concurrency-with-threads/sync2.md)
    - [基于 Send 和 Sync 的线程安全](/markdown/rust/rust-course/advance/concurrency-with-threads/send-sync.md)
  - [全局变量](/markdown/rust/rust-course/advance/global-variable.md)
  - [错误处理](/markdown/rust/rust-course/advance/errors.md)
  - [Unsafe Rust](/markdown/rust/rust-course/advance/unsafe/intro.md)
    - [五种兵器](/markdown/rust/rust-course/advance/unsafe/superpowers.md)
    - [内联汇编](/markdown/rust/rust-course/advance/unsafe/inline-asm.md)
  - [Macro 宏编程](/markdown/rust/rust-course/advance/macro.md)
    <!-- - [SIMD todo](/markdown/rust/rust-course/advance/simd.md) -->
    <!-- - [高阶特征约束(HRTB) todo](/markdown/rust/rust-course/advance/hrtb.md) -->
  - [async/await 异步编程](/markdown/rust/rust-course/advance/async/intro.md)
    - [async 编程入门](/markdown/rust/rust-course/advance/async/getting-started.md)
    - [底层探秘: Future 执行与任务调度](/markdown/rust/rust-course/advance/async/future-excuting.md)
    - [定海神针 Pin 和 Unpin](/markdown/rust/rust-course/advance/async/pin-unpin.md)
    - [async/await 和 Stream 流处理](/markdown/rust/rust-course/advance/async/async-await.md)
    - [同时运行多个 Future](/markdown/rust/rust-course/advance/async/multi-futures-simultaneous.md)
    - [一些疑难问题的解决办法](/markdown/rust/rust-course/advance/async/pain-points-and-workarounds.md)
    - [实践应用：Async Web 服务器](/markdown/rust/rust-course/advance/async/web-server.md)

- [进阶实战1: 实现一个 web 服务器](/markdown/rust/rust-course/advance-practice1/intro.md)
  - [单线程版本](/markdown/rust/rust-course/advance-practice1/web-server.md)
  - [多线程版本](/markdown/rust/rust-course/advance-practice1/multi-threads.md)
  - [优雅关闭和资源清理](/markdown/rust/rust-course/advance-practice1/graceful-shutdown.md)

- [进阶实战2: 实现一个简单 Redis](/markdown/rust/rust-course/advance-practice/intro.md)
  - [tokio 概览](/markdown/rust/rust-course/advance-practice/overview.md)
  - [使用初印象](/markdown/rust/rust-course/advance-practice/getting-startted.md)
  - [创建异步任务](/markdown/rust/rust-course/advance-practice/spawning.md)
  - [共享状态](/markdown/rust/rust-course/advance-practice/shared-state.md)
  - [消息传递](/markdown/rust/rust-course/advance-practice/channels.md)
  - [I/O](/markdown/rust/rust-course/advance-practice/io.md)
  - [解析数据帧](/markdown/rust/rust-course/advance-practice/frame.md)
  - [深入 async](/markdown/rust/rust-course/advance-practice/async.md)
  - [select](/markdown/rust/rust-course/advance-practice/select.md)
  - [类似迭代器的 Stream](/markdown/rust/rust-course/advance-practice/stream.md)
  - [优雅的关闭](/markdown/rust/rust-course/advance-practice/graceful-shutdown.md)
  - [异步跟同步共存](/markdown/rust/rust-course/advance-practice/bridging-with-sync.md)

<!-- -  [Rust 设计模式](/markdown/rust/rust-course/advance-practice/design-pattern.md) -->

- [Rust 难点攻关](/markdown/rust/rust-course/difficulties/intro.md)
  - [切片和切片引用](/markdown/rust/rust-course/difficulties/slice.md)
  - [Eq 和 PartialEq](/markdown/rust/rust-course/difficulties/eq.md)
  - [String、&str 和 str TODO](/markdown/rust/rust-course/difficulties/string.md)
  - [作用域、生命周期和 NLL TODO](/markdown/rust/rust-course/difficulties/lifetime.md)
  - [move、Copy 和 Clone TODO](/markdown/rust/rust-course/difficulties/move-copy.md)
  - [裸指针、引用和智能指针 TODO](/markdown/rust/rust-course/advance/difficulties/pointer.md)

**常用工具链**

---

- [自动化测试](/markdown/rust/rust-course/test/intro.md)

  - [编写测试及控制执行](/markdown/rust/rust-course/test/write-tests.md)
  - [单元测试和集成测试](/markdown/rust/rust-course/test/unit-integration-test.md)
  - [断言 assertion](/markdown/rust/rust-course/test/assertion.md)
  - [用 GitHub Actions 进行持续集成](/markdown/rust/rust-course/test/ci.md)
  - [基准测试 benchmark](/markdown/rust/rust-course/test/benchmark.md)

- [Cargo 使用指南](/markdown/rust/rust-course/cargo/intro.md)
  - [上手使用](/markdown/rust/rust-course/cargo/getting-started.md)
  - [基础指南](/markdown/rust/rust-course/cargo/guide/intro.md)
    - [为何会有 Cargo](/markdown/rust/rust-course/cargo/guide/why-exist.md)
    - [下载并构建 Package](/markdown/rust/rust-course/cargo/guide/download-package.md)
    - [添加依赖](/markdown/rust/rust-course/cargo/guide/dependencies.md)
    - [Package 目录结构](/markdown/rust/rust-course/cargo/guide/package-layout.md)
    - [Cargo.toml vs Cargo.lock](/markdown/rust/rust-course/cargo/guide/cargo-toml-lock.md)
    - [测试和 CI](/markdown/rust/rust-course/cargo/guide/tests-ci.md)
    - [Cargo 缓存](/markdown/rust/rust-course/cargo/guide/cargo-cache.md)
    - [Build 缓存](/markdown/rust/rust-course/cargo/guide/build-cache.md)
  - [进阶指南](/markdown/rust/rust-course/cargo/reference/intro.md)
    - [指定依赖项](/markdown/rust/rust-course/cargo/reference/specify-deps.md)
    - [依赖覆盖](/markdown/rust/rust-course/cargo/reference/deps-overriding.md)
    - [Cargo.toml 清单详解](/markdown/rust/rust-course/cargo/reference/manifest.md)
    - [Cargo Target](/markdown/rust/rust-course/cargo/reference/cargo-target.md)
    - [工作空间 Workspace](/markdown/rust/rust-course/cargo/reference/workspaces.md)
    - [条件编译 Features](/markdown/rust/rust-course/cargo/reference/features/intro.md)
      - [Features 示例](/markdown/rust/rust-course/cargo/reference/features/examples.md)
    - [发布配置 Profile](/markdown/rust/rust-course/cargo/reference/profiles.md)
    - [通过 config.toml 对 Cargo 进行配置](/markdown/rust/rust-course/cargo/reference/configuration.md)
    - [发布到 crates.io](/markdown/rust/rust-course/cargo/reference/publishing-on-crates.io.md)
    - [构建脚本 build.rs](/markdown/rust/rust-course/cargo/reference/build-script/intro.md)
      - [构建脚本示例](/markdown/rust/rust-course/cargo/reference/build-script/examples.md)

**开发实践**

---

- [企业落地实践](/markdown/rust/rust-course/usecases/intro.md)

  - [AWS 为何这么喜欢 Rust?](/markdown/rust/rust-course/usecases/aws-rust.md)

- [日志和监控](/markdown/rust/rust-course/logs/intro.md)
  - [日志详解](/markdown/rust/rust-course/logs/about-log.md)
  - [日志门面 log](/markdown/rust/rust-course/logs/log.md)
  - [使用 tracing 记录日志](/markdown/rust/rust-course/logs/tracing.md)
  - [自定义 tracing 的输出格式](/markdown/rust/rust-course/logs/tracing-logger.md)
  - [监控](/markdown/rust/rust-course/logs/observe/intro.md)
    - [可观测性](/markdown/rust/rust-course/logs/observe/about-observe.md)
    - [分布式追踪](/markdown/rust/rust-course/logs/observe/trace.md)
- [Rust 最佳实践](/markdown/rust/rust-course/practice/intro.md)
  - [日常开发三方库精选](/markdown/rust/rust-course/practice/third-party-libs.md)
  - [命名规范](/markdown/rust/rust-course/practice/naming.md)
  - [面试经验](/markdown/rust/rust-course/practice/interview.md)
  - [代码开发实践 todo](/markdown/rust/rust-course/practice/best-pratice.md)
- [手把手带你实现链表](/markdown/rust/rust-course/too-many-lists/intro.md)
  - [我们到底需不需要链表](/markdown/rust/rust-course/too-many-lists/do-we-need-it.md)
  - [不太优秀的单向链表：栈](/markdown/rust/rust-course/too-many-lists/bad-stack/intro.md)
    - [数据布局](/markdown/rust/rust-course/too-many-lists/bad-stack/layout.md)
    - [基本操作](/markdown/rust/rust-course/too-many-lists/bad-stack/basic-operations.md)
    - [最后实现](/markdown/rust/rust-course/too-many-lists/bad-stack/final-code.md)
  - [还可以的单向链表](/markdown/rust/rust-course/too-many-lists/ok-stack/intro.md)
    - [优化类型定义](/markdown/rust/rust-course/too-many-lists/ok-stack/type-optimizing.md)
    - [定义 Peek 函数](/markdown/rust/rust-course/too-many-lists/ok-stack/peek.md)
    - [IntoIter 和 Iter](/markdown/rust/rust-course/too-many-lists/ok-stack/iter.md)
    - [IterMut 以及完整代码](/markdown/rust/rust-course/too-many-lists/ok-stack/itermut.md)
  - [持久化单向链表](/markdown/rust/rust-course/too-many-lists/persistent-stack/intro.md)
    - [数据布局和基本操作](/markdown/rust/rust-course/too-many-lists/persistent-stack/layout.md)
    - [Drop、Arc 及完整代码](/markdown/rust/rust-course/too-many-lists/persistent-stack/drop-arc.md)
  - [不咋样的双端队列](/markdown/rust/rust-course/too-many-lists/deque/intro.md)
    - [数据布局和基本操作](/markdown/rust/rust-course/too-many-lists/deque/layout.md)
    - [Peek](/markdown/rust/rust-course/too-many-lists/deque/peek.md)
    - [基本操作的对称镜像](/markdown/rust/rust-course/too-many-lists/deque/symmetric.md)
    - [迭代器](/markdown/rust/rust-course/too-many-lists/deque/iterator.md)
    - [最终代码](/markdown/rust/rust-course/too-many-lists/deque/final-code.md)
  - [不错的 unsafe 队列](/markdown/rust/rust-course/too-many-lists/unsafe-queue/intro.md)
    - [数据布局](/markdown/rust/rust-course/too-many-lists/unsafe-queue/layout.md)
    - [基本操作](/markdown/rust/rust-course/too-many-lists/unsafe-queue/basics.md)
    - [Miri](/markdown/rust/rust-course/too-many-lists/unsafe-queue/miri.md)
    - [栈借用](/markdown/rust/rust-course/too-many-lists/unsafe-queue/stacked-borrow.md)
    - [测试栈借用](/markdown/rust/rust-course/too-many-lists/unsafe-queue/testing-stacked-borrow.md)
    - [数据布局 2](/markdown/rust/rust-course/too-many-lists/unsafe-queue/layout2.md)
    - [额外的操作](/markdown/rust/rust-course/too-many-lists/unsafe-queue/extra-junk.md)
    - [最终代码](/markdown/rust/rust-course/too-many-lists/unsafe-queue/final-code.md)
  - [生产级的双向 unsafe 队列](/markdown/rust/rust-course/too-many-lists/production-unsafe-deque/intro.md)
    - [数据布局](/markdown/rust/rust-course/too-many-lists/production-unsafe-deque/layout.md)
    - [型变与子类型](/markdown/rust/rust-course/too-many-lists/production-unsafe-deque/variance-and-phantomData.md)
    - [基础结构](/markdown/rust/rust-course/too-many-lists/production-unsafe-deque/basics.md)
    - [恐慌与安全](/markdown/rust/rust-course/too-many-lists/production-unsafe-deque/drop-and-panic-safety.md)
    - [无聊的组合](/markdown/rust/rust-course/too-many-lists/production-unsafe-deque/boring-combinatorics.md)
    - [其它特征](/markdown/rust/rust-course/too-many-lists/production-unsafe-deque/filling-in-random-bits.md)
    - [测试](/markdown/rust/rust-course/too-many-lists/production-unsafe-deque/testing.md)
    - [Send,Sync和编译测试](/markdown/rust/rust-course/too-many-lists/production-unsafe-deque/send-sync-and-compile-tests.md)
    - [实现游标](/markdown/rust/rust-course/too-many-lists/production-unsafe-deque/implementing-cursors.md)
    - [测试游标](/markdown/rust/rust-course/too-many-lists/production-unsafe-deque/testing-cursors.md)
    - [最终代码](/markdown/rust/rust-course/too-many-lists/production-unsafe-deque/final-code.md)
  - [使用高级技巧实现链表](/markdown/rust/rust-course/too-many-lists/advanced-lists/intro.md)
    - [双单向链表](/markdown/rust/rust-course/too-many-lists/advanced-lists/double-singly.md)
    - [栈上的链表](/markdown/rust/rust-course/too-many-lists/advanced-lists/stack-allocated.md)

**攻克编译错误**

---

- [征服编译错误](/markdown/rust/rust-course/compiler/intro.md)

  - [对抗编译检查](/markdown/rust/rust-course/compiler/fight-with-compiler/intro.md)
    - [生命周期](/markdown/rust/rust-course/compiler/fight-with-compiler/lifetime/intro.md)
      - [生命周期过大-01](/markdown/rust/rust-course/compiler/fight-with-compiler/lifetime/too-long1.md)
      - [生命周期过大-02](/markdown/rust/rust-course/compiler/fight-with-compiler/lifetime/too-long2.md)
      - [循环中的生命周期](/markdown/rust/rust-course/compiler/fight-with-compiler/lifetime/loop.md)
      - [闭包碰到特征对象-01](/markdown/rust/rust-course/compiler/fight-with-compiler/lifetime/closure-with-static.md)
    - [重复借用](/markdown/rust/rust-course/compiler/fight-with-compiler/borrowing/intro.md)
      - [同时在函数内外使用引用](/markdown/rust/rust-course/compiler/fight-with-compiler/borrowing/ref-exist-in-out-fn.md)
      - [智能指针引起的重复借用错误](/markdown/rust/rust-course/compiler/fight-with-compiler/borrowing/borrow-distinct-fields-of-struct.md)
    - [类型未限制(todo)](/markdown/rust/rust-course/compiler/fight-with-compiler/unconstrained.md)
    - [幽灵数据(todo)](/markdown/rust/rust-course/compiler/fight-with-compiler/phantom-data.md)
  - [Rust 常见陷阱](/markdown/rust/rust-course/compiler/pitfalls/index.md)
    - [for 循环中使用外部数组](/markdown/rust/rust-course/compiler/pitfalls/use-vec-in-for.md)
    - [线程类型导致的栈溢出](/markdown/rust/rust-course/compiler/pitfalls/stack-overflow.md)
    - [算术溢出导致的 panic](/markdown/rust/rust-course/compiler/pitfalls/arithmetic-overflow.md)
    - [闭包中奇怪的生命周期](/markdown/rust/rust-course/compiler/pitfalls/closure-with-lifetime.md)
    - [可变变量不可变？](/markdown/rust/rust-course/compiler/pitfalls/the-disabled-mutability.md)
    - [可变借用失败引发的深入思考](/markdown/rust/rust-course/compiler/pitfalls/multiple-mutable-references.md)
    - [不太勤快的迭代器](/markdown/rust/rust-course/compiler/pitfalls/lazy-iterators.md)
    - [奇怪的序列 x..y](/markdown/rust/rust-course/compiler/pitfalls/weird-ranges.md)
    - [无处不在的迭代器](/markdown/rust/rust-course/compiler/pitfalls/iterator-everywhere.md)
    - [线程间传递消息导致主线程无法结束](/markdown/rust/rust-course/compiler/pitfalls/main-with-channel-blocked.md)
    - [警惕 UTF-8 引发的性能隐患](/markdown/rust/rust-course/compiler/pitfalls/utf8-performance.md)

**性能优化**

---

- [Rust 性能优化 todo](/markdown/rust/rust-course/profiling/intro.md)

  - [深入内存 todo](/markdown/rust/rust-course/profiling/memory/intro.md)
    - [指针和引用 todo](/markdown/rust/rust-course/profiling/memory/pointer-ref.md)
    - [未初始化内存 todo](/markdown/rust/rust-course/profiling/memory/uninit.md)
    - [内存分配 todo](/markdown/rust/rust-course/profiling/memory/allocation.md)
    - [内存布局 todo](/markdown/rust/rust-course/profiling/memory/layout.md)
    - [虚拟内存 todo](/markdown/rust/rust-course/profiling/memory/virtual.md)
  - [性能调优 doing](/markdown/rust/rust-course/profiling/performance/intro.md)
    - [字符串操作性能](/markdown/rust/rust-course/profiling/performance/string.md)
    - [深入理解 move](/markdown/rust/rust-course/profiling/performance/deep-into-move.md)
    - [糟糕的提前优化 todo](/markdown/rust/rust-course/profiling/performance/early-optimise.md)
    - [Clone 和 Copy todo](/markdown/rust/rust-course/profiling/performance/clone-copy.md)
    - [减少 Runtime check(todo)](/markdown/rust/rust-course/profiling/performance/runtime-check.md)
    - [CPU 缓存性能优化 todo](/markdown/rust/rust-course/profiling/performance/cpu-cache.md)
    - [计算性能优化 todo](/markdown/rust/rust-course/profiling/performance/calculate.md)
    - [堆和栈 todo](/markdown/rust/rust-course/profiling/performance/heap-stack.md)
    - [内存 allocator todo](/markdown/rust/rust-course/profiling/performance/allocator.md)
    - [常用性能测试工具 todo](/markdown/rust/rust-course/profiling/performance/tools.md)
    - [Enum 内存优化 todo](/markdown/rust/rust-course/profiling/performance/enum.md)
  - [编译优化 todo](/markdown/rust/rust-course/profiling/compiler/intro.md)
    - [LLVM todo](/markdown/rust/rust-course/profiling/compiler/llvm.md)
    - [常见属性标记 todo](/markdown/rust/rust-course/profiling/compiler/attributes.md)
    - [提升编译速度 todo](/markdown/rust/rust-course/profiling/compiler/speed-up.md)
    - [编译器优化 todo](/markdown/rust/rust-course/profiling/compiler/optimization/intro.md)
      - [Option 枚举 todo](/markdown/rust/rust-course/profiling/compiler/optimization/option.md)

<!-- - [标准库解析 todo](/markdown/rust/rust-course/std/intro.md)

  - [标准库使用最佳实践 todo](/markdown/rust/rust-course/std/search.md)
  - [Vector 常用方法 todo](/markdown/rust/rust-course/std/vector.md)
  - [HashMap todo](/markdown/rust/rust-course/std/hashmap.md)
  - [Iterator 常用方法 todo](/markdown/rust/rust-course/std/iterator.md) -->

  <!-- - [配置文件解析 todo](/markdown/rust/rust-course/cookbook/config.md)
  - [编解码 todo](/markdown/rust/rust-course/cookbook/encoding/intro.md)
    - [JSON](/markdown/rust/rust-course/cookbook/encoding/json.md)
    - [CSV](/markdown/rust/rust-course/cookbook/encoding/csv.md)
    - [protobuf](/markdown/rust/rust-course/cookbook/encoding/protobuf.md)
  - [文件系统 todo](/markdown/rust/rust-course/cookbook/file/intro.md)
    - [文件读写](/markdown/rust/rust-course/cookbook/file/file.md)
    - [目录操作](/markdown/rust/rust-course/cookbook/file/dir.md)
  - [网络通信 todo](/markdown/rust/rust-course/cookbook/protocol/intro.md)
    - [HTTP](/markdown/rust/rust-course/cookbook/protocol/http.md)
    - [TCP](/markdown/rust/rust-course/cookbook/protocol/tcp.md)
    - [UDP](/markdown/rust/rust-course/cookbook/protocol/udp.md)
    - [gRPC](/markdown/rust/rust-course/cookbook/protocol/grpc.md)
  - [数据库访问 todo](/markdown/rust/rust-course/cookbook/database.md)
  - [正则表达式 todo](/markdown/rust/rust-course/cookbook/regexp.md)
  - [加密解密 todo](/markdown/rust/rust-course/cookbook/crypto.md)
  - [时间日期](/markdown/rust/rust-course/cookbook/date.md)
  - [开发调试 todo](/markdown/rust/rust-course/cookbook/dev/intro.md)
    - [日志](/markdown/rust/rust-course/cookbook/dev/logs.md)
    - [性能分析](/markdown/rust/rust-course/cookbook/dev/profile.md) -->

<!--
- [Rust区块链入门](/markdown/rust/rust-course/)
- [Rust游戏开发入门](/markdown/rust/rust-course/)
- [Rust前端开发入门](/markdown/rust/rust-course/)
- [Rust和WASM](/markdown/rust/rust-course/) -->

**附录**

---

- [Appendix](/markdown/rust/rust-course/)
  - [关键字](/markdown/rust/rust-course/appendix/keywords.md)
  - [运算符与符号](/markdown/rust/rust-course/appendix/operators.md)
  - [表达式](/markdown/rust/rust-course/appendix/expressions.md)
  - [派生特征 trait](/markdown/rust/rust-course/appendix/derive.md)
  - [prelude 模块 todo](/markdown/rust/rust-course/appendix/prelude.md)
  - [Rust 版本说明](/markdown/rust/rust-course/appendix/rust-version.md)
  - [Rust 历次版本更新解读](/markdown/rust/rust-course/appendix/rust-versions/intro.md)
    - [1.58](/markdown/rust/rust-course/appendix/rust-versions/1.58.md)
    - [1.59](/markdown/rust/rust-course/appendix/rust-versions/1.59.md)
    - [1.60](/markdown/rust/rust-course/appendix/rust-versions/1.60.md)
    - [1.61](/markdown/rust/rust-course/appendix/rust-versions/1.61.md)
    - [1.62](/markdown/rust/rust-course/appendix/rust-versions/1.62.md)
    - [1.63](/markdown/rust/rust-course/appendix/rust-versions/1.63.md)
    - [1.64](/markdown/rust/rust-course/appendix/rust-versions/1.64.md)
    - [1.65](/markdown/rust/rust-course/appendix/rust-versions/1.65.md)
    - [1.66](/markdown/rust/rust-course/appendix/rust-versions/1.66.md)
    - [1.67](/markdown/rust/rust-course/appendix/rust-versions/1.67.md)
    - [1.68](/markdown/rust/rust-course/appendix/rust-versions/1.68.md)
    - [1.69](/markdown/rust/rust-course/appendix/rust-versions/1.69.md)
    - [1.70](/markdown/rust/rust-course/appendix/rust-versions/1.70.md)
    - [1.71](/markdown/rust/rust-course/appendix/rust-versions/1.71.md)
    - [1.72](/markdown/rust/rust-course/appendix/rust-versions/1.72.md)
    - [1.73](/markdown/rust/rust-course/appendix/rust-versions/1.73.md)
    - [1.74](/markdown/rust/rust-course/appendix/rust-versions/1.74.md)
    - [1.75](/markdown/rust/rust-course/appendix/rust-versions/1.75.md)
    - [1.76](/markdown/rust/rust-course/appendix/rust-versions/1.76.md)
    - [1.77](/markdown/rust/rust-course/appendix/rust-versions/1.77.md)
