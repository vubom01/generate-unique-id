Generate Unique Id
====
Follow package: https://github.com/bwmarrin/snowflake

### ID Format
* The ID as a whole is a 63 bit integer stored in an int64
* 39 bits are used to store a timestamp with millisecond precision, using a custom epoch.
* 1 bit is used to distinguish the type of ID.
* 10 bits are used to store a node id - a range from 0 through 1023.
* 13 bits are used to store a sequence number - a range from 0 through 4095.
```
+------------------------------------------------------------------------------+
| 1 Bit Unused | 39 Bit Timestamp | Bit 1 | 10 Bit NodeID | 13 Bit Sequence ID |
+------------------------------------------------------------------------------+
```