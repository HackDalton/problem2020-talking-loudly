# HackDalton: Talking Loudly (Writeup)

> Warning! There are spoilers ahead

You can use the netcat program on your computer to listen, and the grep program to wait for the flag.

```shell
$ nc localhost 8080 | grep flag
part 0 of 6 of the flag is: hackDal
part 1 of 6 of the flag is: ton{gr3
part 2 of 6 of the flag is: p_1s_4_
part 3 of 6 of the flag is: us3ful_
part 4 of 6 of the flag is: t00l_si
part 5 of 6 of the flag is: 1EldtNcQ}
```

As you can see, the flag is `hackDalton{gr3p_1s_4_us3ful_t00l_si1EldtNcQ}`.