package main
import (
	initialize "github.com/haomiao000/DY/server/service/interact/initialize"
	pkg "github.com/haomiao000/DY/server/service/interact/pkg"
	handler "github.com/haomiao000/DY/server/service/interact/handler"
	rpc_interact "github.com/haomiao000/DY/server/grpc_gen/rpc_interact"
	dao "github.com/haomiao000/DY/server/service/interact/dao"
	grpc "google.golang.org/grpc"
	configs "github.com/haomiao000/DY/server/service/interact/configs"
	"net"
	"fmt"
)

func main() {
	db := initialize.InitDB()
	userServ := initialize.InitUser()
	videoServ := initialize.InitVideo()
	grpcServer := grpc.NewServer()
	impl := &handler.InteractServiceImpl{
		FavoriteMysqlManager: dao.NewMysqlManager(db),
		CommentMysqlManager: dao.NewMysqlManager(db),
		UserManager: pkg.NewUserClient(userServ),
		VideoManager: pkg.NewVideoClient(videoServ),
	}
	rpc_interact.RegisterInteractServiceImplServer(grpcServer , impl)

	listener, err := net.Listen("tcp", configs.InteractServerAddr)
	if err != nil {
		fmt.Printf("Failed To Listen On Addr %s: %v", configs.InteractServerAddr,err)
	}
	if err := grpcServer.Serve(listener); err != nil {
		fmt.Printf("Failed To Serve GRPC Server: %v", err)
	}
}