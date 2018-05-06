# GoCryptoData

Check if random ETH wallets have a balance.

The possibility that this program chooses a wallet with a balance is almost zero (64^36 possibilities).
But when it does, it saves the wallet address, private key and the balance into a file balance.txt.

To make this work you have to compile the ethereum-go client first.
(https://github.com/ethereum/go-ethereum/wiki/Developers'-Guide)

To do this you need to have a gcc in the classpath (e.g. tdm64-gcc-5.1.0-2.exe for windows).
