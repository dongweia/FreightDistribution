package wsconn

import "sync"
//维护一个并发安全的连接池
var SM sync.Map
