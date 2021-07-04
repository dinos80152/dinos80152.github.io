# AWS Docker Container CPU Utilization Running at High Percentage

## Problem

Our service built on AWS EC2 in docker container getting high cpu utilization recently.

![container cpu utilization](https://farm5.staticflickr.com/4305/35589776790_ca7891c347_b.jpg)

We thought it is some resource leaking in our process, so we reviewed our code and refactored.

However, it still happen twice a day.

## Solution

Finally we found there is a limit of cpu usage in EC2 called **"CPU Credits"**.

EC2 consumes CPU Credits every 15 minutes, after credits exhausted, **EC2 CPU Ulitization will be blocked at 10%** on a t2.micro instance.

![instance cpu credit balance](https://farm5.staticflickr.com/4320/35692577110_d18208861b_b.jpg)

![instance cpu utilization](https://farm5.staticflickr.com/4325/35241900424_e9cb5acabb_b.jpg)

After CPU Utilization has been blocked, there is not enough CPU for container usage, so container CPU goes over 100%.

## Estimation

for **t2.micro**, the average CPU utilization for 24 hours is **12% CPU** utilization

```math
30 credits + 6 earn credits per hour * 24 hours = 174 credits a day
174 credits / 24 hours = 7.25 credits per hour
7.25 credits * 10 CPU utilization / 6 credits = 12% CPU utilization
```

## Reference

* [CPU Credits@AWS User Guide for Linux Instances](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/t2-instances.html#t2-instances-cpu-credits)