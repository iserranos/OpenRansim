# OpenRansim

OpenRansim is an Open Source tool able to check the level of protection 
that a computer has over a Ransomware attack.

By simulating 10 differents cases of Ransomware infections, you could check if your systems is well protected. 
Under a controlled environment, OpenRansim launches a set of more common Ransomware scenarios and will give feedback on how the system is prepared from such kind of attacks.

## Install

OpenRansim has been developed in Golang, so it is needed for compiling and building the executable binary. Look at [this](https://golang.org/doc/install "Golang installation page") page to install it.

Once Golang is installed on the system, the Makefile file specifies some tasks, such as install build, format and clean, to help you work with the tool.

## Scenarios

As mentioned previously, OpenRansim runs 10 different scenarios with the aim of simulating a Ransomware infection. You can try all these:

   * InsideCryptor: Encrypts the data and overwrites the original files.
  
 
   * LockyVariant - Simulates one of the countless variables of the Locky ransomware.
   * Mover - Encrypt the data in a folder other than the original and delete the original files.
   * Replacer - Replaces the contents of the original files.
   * Streamer - Encrypts all data and groups them into a single file.
   * StrongCryptor - Encrypt the data and delete the original files safely.
   * StrongCryptorFast - Encrypts the data and deletes the original files.
   * StrongCrytptorNet - Encrypts data, clears original files and simulates an HTTP connection.
   * ThorVariant - Simulates one of the countless variables of ransomware Thor.
   * WeakCryptor - Uses weak encryption to encrypt data and removes original files.