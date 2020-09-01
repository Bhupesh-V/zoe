# zoe

> A simple transactional key:value store for learning purpose

The shell accepts the following commands:

- `SET [key] [value]`: Sets the given key to the specified value. If the key is already present, overwrite the old value.
- `GET [key]`: Prints out the current value of the specified key. If the key has not been set, it prints a default message.
- `DELETE [key]`: Deletes the given key. If the key has not been set, ignore.
- `COUNT [value]`: Returns the number of keys that have been set to the specified value. If no keys have been set to that value, prints 0.
- `BEGIN`: Starts a transaction. These transactions allow you to modify the state of the system and commit or rollback your changes.
- `END`: Ends a transaction, everything done within the "active" transaction is lost.
- `ROLLBACK`: Throws away changes made within the context of the active transaction and ends the active transaction. If no transaction is active, prints NO TRANSACTION
<!-- - `COMMIT`: Commits the changes made within the context of the active transaction and ends the active transaction.-->