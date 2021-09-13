## Testing the application

To test the application make sure the adb server is running by using one of the following way:

`adb server` / `adb devices`

Once adb server is up and running, run the Go application 

`go run ./main.go`

the output will be shown as follows

```
Number of byte  16
Response result :  OKAY
Real response value : 0029
Number of byte  22
Response result :  OKAY
Number of byte  18
Response result :  OKAY
Real response value : Linux localhost 3.18.91+ #1 SMP PREEMPT Tue Jan 9 20:30:51 UTC 2018 x86_64
```

More information about the `adb` protocol can be found inside Android source code:

https://android.googlesource.com/platform/system/adb/+/refs/heads/master/OVERVIEW.TXT
https://android.googlesource.com/platform/system/adb/+/refs/heads/master/SERVICES.TXT
https://android.googlesource.com/platform/system/adb/+/refs/heads/master/SYNC.TXT