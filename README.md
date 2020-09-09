# zoe

> A simple transactional key:value store for educational purposes



### [🖱️ Learn how I made `zoe` from scratch](https://www.freecodecamp.org/news/design-a-key-value-store-in-go/)



The shell accepts the following commands:

|   Command  |                                                             Description                                                            |
|:----------:|:----------------------------------------------------------------------------------------------------------------------------------:|
|    `SET`   | Sets the given key to the specified value. A key can also be updated.                                                              |
|    `GET`   | Prints out the current value of the specified key.                                                                                 |
|   `DELETE` | Deletes the given key. If the key has not been set, ignore.                                                                        |
|   `COUNT`  | Returns the number of keys that have been set to the specified value. If no keys have been set to that value, prints 0.            |
|   `BEGIN`  | Starts a transaction. These transactions allow you to modify the state of the system and commit or rollback your changes.          |
|    `END`   | Ends a transaction, everything done within the "active" transaction is lost.                                                       |
| `ROLLBACK` | Throws away changes made within the context of the active transaction. If no transaction is active, prints "No Active Transaction" |
|  `COMMIT`  | Commits the changes made within the context of the active transaction and ends the active transaction.                             |

## Demo

![zoe-demo-3](https://user-images.githubusercontent.com/34342551/92362469-aa2a7700-f10d-11ea-8426-1e8462b66d18.gif)

## Usage

If you just wanna play, ⬇ Download the build from [releases](https://github.com/Bhupesh-V/zoe/releases)

## License

Copyright © 2020 [Bhupesh Varshney](https://github.com/Bhupesh-V).<br />
This project is [MIT](https://github.com/Bhupesh-V/zoe/blob/master/LICENSE) licensed.
