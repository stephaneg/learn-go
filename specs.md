
# common standards
uint32 = 4 octets
uint8 = 1 octet = 8 bits = 2^8 nombres = 256 caractères. Soit en base Hexadecimale : 16 * 16 soit 0xab

## hash
Hash : double SHA-256 hash

## Merkle Trees

## signatures

## adresses
To define in this context (Actor Adress)
* version (1, uint8) : 1 byte of 0, on the test network this version is 1 byte of 1111
* key_hash : version concatenated with RIPEMD-160(SHA-256(public key))
* checksum : 1st 4 bytes of SHA-256(SHA-256(key_hash))
* Bitcoin Address ** TO BE REPLACED IN THIS CONTEXT ** Base58Encode(key_hash concatenated with checksum)

# commons structures

## message structure (field size in bytes)
* magic (4, uint32) : magic data indicating message origin network. main	0xD9B4BEF9 is sent over wire as F9 BE B4 D9
* command (12, char) : ascii string identifying the packet content
* length (4, uint32) : length of payload in # of bytes
* payload (?, byte) : actual data
* checksum (4, uint32) : first 4 bytes of sha256(sha256(payload))




## network adress
* time (4, uint32)
* services (8, uint64) : same services listed in version
* IPV64 (16, char) : IPV6 / IPV4 adress
* port (2, uint16) : port number


## block headers
todo




# message type

## version
payload :
* version (4, uint2)
* services (8, uint64) : bitfield of features to be enabled for this connection
* addr_recv : net_adress
* addr_from : net_adress


## verack

## addr
* count (?, var_int) : number of adresses
* addr_list

## inv
* count(?, var_int) : number of inventory objects
* inventory (inv_vect[]) : inventory vectors
