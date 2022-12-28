package main

import (
	"bufio"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
	"os"
	tgpb "tgbot/pb"
)

func Scan() string {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	if err := in.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
	}
	return in.Text()
}

func main() {

	conn, err := grpc.Dial("localhost:5300", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}

	var word string

	for {
		fmt.Println("Введите слово")
		word = Scan()

		defer conn.Close()
		client := tgpb.NewTGClient(conn)
		request := &tgpb.Message{
			Message: word,
		}

		response, err := client.Echo(context.Background(), request)

		if err != nil {
			grpclog.Fatalf("fail to dial: %v", err)
		}

		fmt.Println(response.Message)
	}
}
