### FTX
Basic full-text indexing and search library.

I wrote this for learning. It is not tweaked for performance, but should work fine for small datasets.

#### Limitations
Indexes are in-memory only. You'll have to implement your own storage layer if you need it.

#### Todo
* Query parser
* More filters & tokenizers
* More indexes, esp ranking ones (tf-idf)