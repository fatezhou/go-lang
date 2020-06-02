package socket

import (
	"errors"
	"fmt"
	"net"
	"strings"
)

type SocketServer struct{
	listen net.Listener
	callback func(*SocketConn)
}

const(
	EVENT_READ = 0
	EVENT_DISCONNECT = 1
	_log_debug = 0
	_log_info = 1
	_log_warn = 2
	_log_err = 3
)

type SocketConn struct{
	net.Conn
	Buffer []byte
	Event int
}

var printer func(format string, a ...interface{})(n int, err error)

func defaultCallback(conn *SocketConn){

}

func (s *SocketConn)Send(buffer []byte)(n int, err error){
	return s.Write(buffer)
}

func (s *SocketConn)Close(){
	s.Close()
}

func _Print(lv int, format string, a ...interface{}){
	if printer == nil{
		fmt.Printf(format, a...)
	}
}

func init(){
	printer = nil
}

func NewSocketServer()*SocketServer{
	return &SocketServer{
		callback:defaultCallback,
	}
}

func (s *SocketServer)SetCallback(ptr func(conn *SocketConn)){
	s.callback = ptr
}

func SetPrinter(ptr func(format string, a ...interface{})(n int, err error)){
	printer = ptr
}

func (s *SocketServer)Listen(ip, protocol string)(err error){
	err = nil
	if strings.ToLower(protocol) != "tcp" && strings.ToLower(protocol)!= "udp"{
		err = errors.New("protocol error, must be tcp or udp")
		_Print(_log_debug,"%+v", err)
		return
	}

	s.listen, err = net.Listen(protocol, ip)
	if err != nil{
		_Print(_log_debug,"%+v", err)
		return
	}
	defer s.listen.Close()
	_Print(_log_info,"wait for clients")
	for{
		conn, _err := s.listen.Accept()
		if _err != nil{
			_Print(_log_err,"%+v", _err)
			continue
		}

		_Print(_log_debug,"client[%s] in", conn.RemoteAddr().String())

		socketConn := &SocketConn{}
		socketConn.Conn = conn
		socketConn.Buffer = make([]byte, 2048)
		//socketConn.SetReadDeadline(time.Now().Add(5 * time.Second))
		go s.HandleConn(socketConn)
	}
}

func (s *SocketServer)HandleConn(socketConn *SocketConn){
	for{
		n, err := socketConn.Conn.Read(socketConn.Buffer)
		if err != nil{
			_Print(_log_err,"client[%+v] err:%+v", socketConn.RemoteAddr(), err)
			socketConn.Event = EVENT_DISCONNECT
			s.callback(socketConn)
			socketConn.Close()
			return
		}else{
			_Print(_log_debug, "recv:[%d]", n)
			socketConn.Event = EVENT_READ
			s.callback(socketConn)
		}
	}
}