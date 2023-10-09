
### CAT 原理
     The cat utility reads files sequentially, writing them to the standard
     output.  The file operands are processed in command-line order.  If file
     is a single dash (`-') or absent, cat reads from the standard input.  If
     file is a UNIX domain socket, cat connects to it and then reads it until
     EOF.  This complements the UNIX domain binding capability available in
     inetd(8).



     The command:

           cat file1 file2 > file3

     will sequentially print the contents of file1 and file2 to the file
     file3, truncating file3 if it already exists.  See the manual page for
     your shell (i.e., sh(1)) for more information on redirection.
