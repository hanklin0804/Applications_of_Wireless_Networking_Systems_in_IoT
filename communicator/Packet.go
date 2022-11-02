package communicator

type MyPacket struct {
	ID       string
	Content  string
	Response string
	Status   int
}

/*
0 default
1 sending
2 retry
3 done
*/
