# guard-trader
Bot trader for exchanges operations

                                                    +---------------+
                                       +------------+    KeyStore   |
                                       v            |               |
                                                    +---------------+
                               +----------------+
                +---+account+--+    Manager     +--------------+
                |              |                |              |
                |              +----+-----------+              |
                |                   |           |              |
            +----+-----+        +----+----+   +--+---+      +---+----+
        +---+ binance1 |        | bitfinex|   |kraken|      |binance2+-----+
        |   +-------+--+        +----+----+   +------+      +--------+     |
        v           v                v                                     v
    +----+-----+  +--+--+        +----+-----+                        +------+---+
    |operation1|  | OP2 |        |operation1|                        |operation1|
    +----------+  +-----+        +----------+                        +----------+


### Tasks list:
- [x] Architeture to support N exchanges acc
- [x] Code structure of project
- [x] Exchange interface contract for manager
- [x] Manager create new exchange
- [x] Repository Badger
- [x] Tests for connections and req/resp
- [x] Basic Cli commands
- [ ] Implement validations and errors

### Backlog ideas to split in tasks:
- [ ] KeyStore accounts
- [ ] Operations: flow for generate analisys and orders
- [ ] Operations: stop limit buy/sell
- [ ] Guard bot skills struct
- [ ] Graphical analisys skills(indicators)
- [ ] Manual Bot decisions
- [ ] Automated Bot decisions


