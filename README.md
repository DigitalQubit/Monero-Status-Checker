# Monero-Status-Checker

Terminal GUI Application for keeping track of mining statuses from various popular pools.

This application can be used to monitor the hashrate, the percentage of valid shares, amount of monero due to be paid, amount of monero already paid, and provide a hashrate history chart **directly to the console!** No more having to go to the website, and dealing with the hassle of clicking through tabs, etc. 

## Supported pools: 

1. [supportxmr](https://supportxmr.com)
2. [xmrpool](https://xmrpool.net)
3. [viaxmr](https://viaxmr.com)
4. [hashvault](https://monero.hashvault.pro)
5. [moriaxmr](https://moriaxmr.com)
6. [moneroocean](https://moneroocean.stream)

## Features 

1. The console can be dynamically resized, and the gui will maintain its original format
**Note: Sorry for the low gif quality, but you get the point.**
![Demo Gif](https://github.com/DigitalQubit/Monero-Status-Checker/raw/master/mydemo.gif)

2. The Valid Shares percentage bar color will turn red if the percentage of Valid Shares drops below 98%

3. The Hashrate color will change based on current hashrate

4. The Hashrate History Graph dots will also change color depending on the current hashrate

![Demo Image](https://github.com/DigitalQubit/Monero-Status-Checker/raw/master/DemoCheckStats.png)

## Usage

### Windows
`CheckStats.exe <Your Monero Address Here> <The Pool # Here>`

### Linux 
`./CheckStats <Your Monero Address Here> <The Pool # Here>`

### Example 
**Fake Public Address, and checking stats from supportxmr!**
`./CheckStats 4hd54xjkhsdjf4jkhsdf4jzukh84jkhsdu8edDedf55sdkljfa8asdfa 1`

### What are the monero pool numbers?
1 == [supportxmr](https://supportxmr.com)
2 == [xmrpool](https://xmrpool.net)
3 == [viaxmr](https://viaxmr.com)
4 == [hashvault](https://monero.hashvault.pro)
5 == [moriaxmr](https://moriaxmr.com)
6 == [moneroocean](https://moneroocean.stream)

## Please sir, another coffee?

**XMR:** 439Jr3ATzETf7ARQgmqAybW9B4htJi5DmDU97ZFffMDXAinbXPCAbydf8Zy1ELqvFV59JYQkn2zswMmt6S2PbUajRZ4BdEQ
**BTC:** 17tAcYQPD44QHqCXeByYczDeqV8MPnmAwq
