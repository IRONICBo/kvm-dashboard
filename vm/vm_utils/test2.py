from paramiko import SSHClient, AutoAddPolicy
import time
import re
import os
import sys
import progressbar
import stat

def create_connection(host):
    # username = 'asklv'
    # password = ''
    port = 22
    connection_info = {
        'port': port,
        # 'username': username,
        # 'password': password
    }

    client = SSHClient()
    client.set_missing_host_key_policy(AutoAddPolicy())
    client.connect(host, timeout=10, auth_timeout=10, **connection_info)
    ssh_session = client.invoke_shell()
    ssh_session.settimeout(10.0)

    return ssh_session


def send_commands(ssh_session, commands, sleep_time=1):
    for command in commands:
        ssh_session.send(command)
        time.sleep(sleep_time)

    output = ssh_session.recv(1048576)
    decoded_output = output.decode()
    return decoded_output


console_commands = [
        'virsh console vw-vta\n',
        '\n',
        '\n',  # Place the username of the VTA here
        '\n'  # Place the password of the VTA here
    ]
show_mgmt_commands = [
        'ifconfig ens2 | grep Bcast\n'
    ]

exit_console = [
        'exit\n'
        '\x1b'
    ]
validate_commands = [
        'virsh list\n'
    ]


def validation():
    host = input('What is the IP address? ')
    print('\n')
    print(f'\033[1;33m--< Checking {host} for a valid VTA >------------\033[0m')

    try:
        ssh_session = create_connection(host)
    except Exception as l:
        print(f"\033[1;31mCannot connect to {host}!\033[0m")
        return

    validate = send_commands(ssh_session, validate_commands)

    if 'y1564' in validate:
        print(f"\033[1;32mThere is a Y1564 VTA running! Obtaining information...\033[0m")
        ssh_session = create_connection(host)
        console = send_commands(ssh_session, console_commands)

        try:
            show_mgmt = send_commands(ssh_session, show_mgmt_commands, sleep_time=2)

        except Exception as e:
            print(f"\033[1;31mCouldn't reach the console on " f"\033[1;33m{host}\033[0m"f"\033[1;31m. This VTA will need to be rebuilt.\033[0m")



        if 'Login incorrect' in show_mgmt:
            print(f"\033[1;31m--< Begin ERROR MESSAGE >------------\033[0m")
            print(show_mgmt)
            print(f"\033[1;31m--< End ERROR MESSAGE >------------\033[0m")
            print(f"\033[1;31mThis VTA has the incorrect password!\033[0m")
            print(f'{host},VTA Password Error', file=f)
            exit = send_commands(ssh_session, exit_console)
            return
        else:
            try:
                mgmt = show_mgmt.split('addr:')[1].split(" ")[0]
            except Exception as g:
                print(f"\033[1;31mThis VTA is corrupt and will need to be rebuilt!\033[0m")
                exit = send_commands(ssh_session, exit_console)
                return
            print("Y1564 VTA IP: ", mgmt)
            exit = send_commands(ssh_session, exit_console)

    else:
        print("\033[1;33mThere is NOT a Y1564 VTA running on \033[0m"f"\033[1;34m {host}\033[0m")
    
    ssh_session.close()

if __name__ == '__main__':
    full_check()