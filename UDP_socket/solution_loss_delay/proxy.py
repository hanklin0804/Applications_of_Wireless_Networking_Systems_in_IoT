import socket
import random
import time
import select
from threading import Thread



localIP     = "127.0.0.1"
localPort   = 5406

# Create a datagram socket
UDPProxySocket = socket.socket(family=socket.AF_INET, type=socket.SOCK_DGRAM)
# Bind to address and ip
UDPProxySocket.bind((localIP, localPort))

print("UDP Proxy up and listening")


flag = 0
def delay(t):
    time.sleep(t)

    # Sending a msg to server
    UDPProxySocket.sendto(proxyMsg, ("127.0.0.1", 5405))
    print(123)

# # threading define
delay100ms = Thread(target = delay  , args = (0.1,))


# Listen for incoming datagrams
drop_num = 0
deley_num = 0
while(True):
    # try: 
        # timeout setting 5ms
        ready = select.select([UDPProxySocket], [], [], 0.002)
        if ready[0]:
            proxyMsg , proxyIP = UDPProxySocket.recvfrom(1024)
            # print("Message from Client: ",proxyMsg.decode())
            # print("Client IP Address: ",proxyIP)

            # it drops each received packet with 10% probability if i is even number. 
            if int(proxyMsg.decode().split(' ')[1]) % 2 == 0:
                if not random.randint(0,9) == 0:
                    # Sending a msg to server
                    UDPProxySocket.sendto(proxyMsg, ("127.0.0.1", 5405))
                else: 
                    print(proxyMsg.decode(),'is dropped')
                    drop_num += 1 

            # it delays 100 ms the received packet with 5% probability before forwarding to the server if i is odd number. 
            else:
                if random.randint(0,19) == 0: 
                    deley_num += 1
                    print(proxyMsg.decode(),' id delays 100 ms ')
                    if not delay100ms.is_alive():
                        delay100ms.start()
                    delay100ms.join()
                else:
                    # Sending a msg to server
                    UDPProxySocket.sendto(proxyMsg, ("127.0.0.1", 5405))
                    
                

            if int(proxyMsg.decode().split(' ')[1]) >= 10000:
                print('num of drop: ',drop_num)
                print('num of delay: ',deley_num)
                break

    # except:
    #     # Create a datagram socket
    #     UDPProxySocket = socket.socket(family=socket.AF_INET, type=socket.SOCK_DGRAM)

    #     # Bind to address and ip
    #     UDPProxySocket.bind((localIP, localPort))

        
