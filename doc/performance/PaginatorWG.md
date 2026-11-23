---
title: 性能测试（PaginatorWG）
description: PaginatorWG 协程分页相对旧版 Paginator 在真实生产与远距离环境下的延迟降低数据。
order: 10
category: 性能
---

# 性能测试 Performance Test

## 1. 实际生产环境延迟降低：约 71.98%

（Latency Decrease 71%）

阿里云 RDS 项目实测：

- 原版：![img_1.png](img_1.png)
- 新版：![img.png](img.png)

## 2. 远距离环境延迟减少：约 28.91%

以下内容仅统计程序处理时间（Latency only for the binary file）：

| Old-Paginator | New-PaginatorWG |
|---------------|-----------------|
| 1144ms        | 971.641152ms    |
| 1134ms        | 968.76604ms     |
| 1140ms        | 964.523454ms    |
| 1137ms        | 968.157527ms    |
| 1142ms        | 882.433464ms    |
| 1143ms        | 974.042662ms    |
| 1138ms        | 970.14531ms     |
| 1146ms        | 976.033914ms    |
| 1144ms        | 973.388182ms    |
| 1138ms        | 968.376514ms    |
| 1150ms        | 975.206023ms    |

- 极限性能提高：28.9115646258503%
- 最低性能提高：0.17827868852459%