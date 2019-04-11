# The GOCH

Sometimes bad things happen, for example You work in corportion.
And You don't have telnet, or You're just too lazy to write bash loop.
The solution is **GOCH**. WHY?

- Sample
- Small
- Asynchronous
- Easy to configure, just add next line to list.json

  ```json
   {
     "ip": "10.100.123.10",
     "port": 443
   },
   ```

- Easy to build(only standard libs)

   ```bash
   go build goch.go
   ```

- Easy to run

  ```bash
   chmod +x goch
   ./goch
  ```

- Readable output

  ```console
  #################################

  Successful connected to: 10.100.123.10:443

  dial tcp 10.100.123.10:444: connect: connection refused

  dial tcp 10.100.123.10:445 i/o timeout

  #################################
  ```
