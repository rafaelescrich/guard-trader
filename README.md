# guard-trader
Bot trader for exchanges operations

                                                       +---------------+
                                          +------------+    Account    |
                                          v            |               |
                                                       +---------------+
                                  +----------------+
                   +---+account+--+    Manager     +--------------+
                   |              |                |              |
                   |              +----+-----------+              |
                   |                   |           |              |
              +----+-----+        +----+----+   +--+---+      +---+----+
          +---+ binance1 |        | bitfinex|   |kraken|      |binance2+-----+N
          |   +-------+--+        +----+----+   +------+      +--------+     |
          v           v                v                                     v
     +----+-----+  +--+--+        +----+-----+                        +------+---+
+----+operation1|  | OP2 |        |operation1|                        |operation1|
|    +----------+  +-----+        +----------+                        +----------+
+->   Analisys
|-->  Orders
+---> objectives


###Tasks list:
- [x] Flow architeture to support N exchanges acc lvl 1
- [x] Code structure of project
- [x] Exchange interface contract for manager
- [x] Manager create new exchange
- [x] Accounts lvl 1
- [x] Repository BoltDB lvl 1
- [x] Operations struct lvl 1
- [ ] Tests for connections and req/resp
- [ ] Implement validations and errors


###Backlog ideas to split in tasks:
- [ ] Operations: flow for generate analisys and orders
- [ ] Operations: stop limit buy/sell
- [ ] Cli commands
- [ ] Guard bot skills struct
- [ ] Graphical analisys skills(indicators)
- [ ] Manual Bot decisions
- [ ] Automated Bot decisions


