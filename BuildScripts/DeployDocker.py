import call_cmd

if __name__ == "__main__":
    call_cmd.call('docker rm WizzHomePage -f')
    call_cmd.call('docker build -t wizz .')
    call_cmd.call('docker run --name WizzHomePage -d -p 8080:8080 wizz')