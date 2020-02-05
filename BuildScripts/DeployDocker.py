import call_cmd

if __name__ == "__main__":
    call_cmd.call('docker rm WizzHomePage -f')
    call_cmd.call('docker build -t wizz-home-page .')
    call_cmd.call('docker run --name WizzHomePage -d -p 8080:8080 -v /var/www/data:/go/release/Wizz-Home-Page/data wizz-home-page')