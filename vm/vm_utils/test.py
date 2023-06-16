import subprocess
import pty
import time

def connect_to_vm_console(vm_uuid):
    # use -tt to solve these problem
    cmd = ["ssh", "-tt", "asklv@localhost"]
    process = subprocess.Popen(cmd, stdin=subprocess.PIPE, stdout=subprocess.PIPE)
    time.sleep(5)
    print("-1 ok")

    output, _ = process.communicate()
    output = output.decode('utf-8')  # 将字节串转换为字符串
    print(output)
    # 输入用户名和密码（根据需要自行更改）
    process.stdin.write(b"sudo virsh console --force 33ccd0be-a3a0-48ba-a5f7-b8d297cc7048\n")
    # process.stdin.write(b"sudo virsh console --force 33ccd0be-a3a0-48ba-a5f7-b8d297cc7048\n")
    time.sleep(5)
    print("0 ok")
    process.stdin.write(b'\n')
    time.sleep(5)
    print("1 ok")
    process.stdin.write(b'root\n')
    time.sleep(5)
    print("2 ok")
    process.stdin.write(b'Ww1270278575\n')
    time.sleep(5)
    print("3 ok")
    process.stdin.write(b'top\n')
    time.sleep(3)
    print("4 ok")
    process.stdin.write(b'exit\n')
    time.sleep(3)
    print("5 ok")
    process.stdin.write(b'\x01D')
    time.sleep(3)
    print("6 ok")
    process.stdin.flush()

    time.sleep(1)

    # 读取控制台输出
    # output = process.stdout.read()
    # print(output)
    # print(output.decode("utf-8")) # 可以直接解码的

    # 关闭输入流和进程
    process.stdin.close()
    # process.wait()

# 虚拟机的 UUID 或名称（根据实际情况自行更改）
vm_uuid = "33ccd0be-a3a0-48ba-a5f7-b8d297cc7048"
# vm_uuid = "vm1"
connect_to_vm_console(vm_uuid)
