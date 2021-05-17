rem A给B签名
rem private_key：A的私钥
rem chequebook:  A的合约地址
rem beneficiary: B的地址
rem payout: token数量，单位wei


./signer.exe -private_key=4a3257d745cd03a02293f774c781342bd063f4eafde6575c1f22560dbb8eeee5  -chequebook=0x4BF21a02fB505Ce8335FD66430Dc50C31fD677c0  -beneficiary=0x902Dda3CB9281e44B974f825980a875E056682A3  -payout=10000

rem 签名后可以调用A合约， https://goerli.etherscan.io/address/0x4BF21a02fB505Ce8335FD66430Dc50C31fD677c0#writeContract
cashChequeBeneficiary

接收gBzz的合约地址：  0xFE1FC82d87407b87e698C1629822Ecd95bB2BBf9
数量： 10000
签名： 0xea6c2cb80e4f82fd454bf58e94dc8096538897e7a70a7fb993e792b4b24ed6146360549b09ecfb94c0cc0e3f6cbbd7f0409e9c32a984eeb88feaa039324a73721b
