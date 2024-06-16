package geecache

// 一个接口，实现后可以根据PickPeer方法获得缓存客户端
type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}

// 一个客户端的接口，有了客户端后就可以使用Get方法获取缓存
type PeerGetter interface {
	Get(group string, key string) ([]byte, error)
}
