package gcache

import (
	"fmt"
	"github.com/leango/gcache/pb"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"github.com/leango/gcache/consistenthash"
)

const (
	defaultBasePath = "/_geecache/"
	defaultReplicas = 5
)

type HTTPPool struct {
	// this peer's base URL, e.g. "https://example.net:8000"
	self     string
	basePath string

	mu          sync.Mutex // guards peers and httpGetters
	peers       *consistenthash.Map
	httpGerrers map[string]*httpGetter // keyed by e.g. "http://10.0.0.2:8008"
}

func NewHTTPPool(self string) *HTTPPool {
	return &HTTPPool{
		self:     self,
		basePath: defaultBasePath,
	}
}

func (p *HTTPPool) Log(format string, v ...interface{}) {
	log.Printf("[Server %s] %s", p.self, fmt.Sprintf(format, v...))
}

func (p *HTTPPool) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.Log("recv req :%s from %s", r.URL.Path, r.RemoteAddr)
	if !strings.HasPrefix(r.URL.Path, p.basePath) {
		panic("HTTPPool serving unexpected path: " + r.URL.Path)
	}
	p.Log("%s %s", r.Method, r.URL.Path)
	// /<basepath>/<groupname>/<key> required
	parts := strings.SplitN(r.URL.Path[len(p.basePath):], "/", 2)
	if len(parts) != 2 {
		http.Error(w, "bad reqeust", http.StatusBadRequest)
	}
	groupName := parts[0]

	key := parts[1]
	group := GetGroup(groupName)

	if group == nil {
		http.Error(w, "no such group"+groupName, http.StatusNotFound)
		p.Log("no such group %s", group.name)
		return
	}

	view, err := group.Get(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		p.Log("group get err:%v", err)
		return
	}
	body ,err :=proto.Marshal(&pb.Response{Value: view.ByteSlice()})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}


	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(body)
}

type httpGetter struct {
	baseURL string
}

func (h *httpGetter) Get(in *pb.Request,out *pb.Response)  error {
	//u := fmt.Sprintf(
	//	"%v%v/%v",
	//	h.baseURL,
	//	url.QueryEscape(group),
	//	url.QueryEscape(key),
	//)
	//res, err := http.Get(u)
	//if err != nil {
	//	return nil, err
	//}
	//defer res.Body.Close()
	//
	//if res.StatusCode != http.StatusOK {
	//	return nil, fmt.Errorf("server returned: %v", res.Status)
	//}
	//
	//bytes, err := ioutil.ReadAll(res.Body)
	//if err != nil {
	//	return nil, fmt.Errorf("reading response body: %v", err)
	//}
	//
	//return bytes, nil

	u := fmt.Sprintf(
		"%v%v/%v",
		h.baseURL,
		url.QueryEscape(in.GetGroup()),
		url.QueryEscape(in.GetKey()),
	)
	res, err := http.Get(u)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("server returned: %v", res.Status)
	}

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return  fmt.Errorf("reading response body: %v", err)
	}
	if err = proto.Unmarshal(bytes,out);err !=nil {
		return fmt.Errorf("decoding response body: %v", err)
	}
	return nil
}

//???????????????
var _ PeerGetter = (*httpGetter)(nil)

//PeerPicker

func (p *HTTPPool) Set(peers ...string) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.peers = consistenthash.New(defaultReplicas, nil)
	p.peers.Add(peers...)
	//p.Log("%s set peers:%v", p.self, peers)
	p.httpGerrers = make(map[string]*httpGetter, len(peers))

	for _, peer := range peers {
		p.httpGerrers[peer] = &httpGetter{baseURL: peer + p.basePath}
	}
}

// PickPeer picks a peer according to key

func (p *HTTPPool) PickPeer(key string) (PeerGetter, bool) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.Log("start pick key:%s", key)
	var peer string
	if peer = p.peers.Get(key); peer != "" && peer != p.self {
		p.Log("Pick peer %s", peer)
		return p.httpGerrers[peer], true
	}
	p.Log("Pick peer failt")
	return nil, false
}

var _ PeerPicker = (*HTTPPool)(nil)
