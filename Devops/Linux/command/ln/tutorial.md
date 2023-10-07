# Ln 命令

ln 命令的作用。 什么是 【软链接】 【硬链接】 The ln utility creates a new directory entry (linked file) which has the same modes as the original file.  It is useful for
     maintaining multiple copies of a file in many places at once without using up storage for the “copies”; instead, a link
     “points” to the original copy.  There are two types of links; hard links and symbolic links.  How a link “points” to a file is
     one of the differences between a hard and symbolic link1、 By default, ln makes hard links.  A hard link to a file is indistinguishable from the original directory entry; any changes to
     a file are effectively independent of the name used to reference the file.  Hard links may not normally refer to directories
     and may not span file systems.2、A symbolic link contains the name of the file to which it is linked.  The referenced file is used when an open(2) operation is
     performed on the link.  A stat(2) on a symbolic link will return the linked-to file; an lstat(2) must be done to obtain
     information about the link.  The readlink(2) call may be used to read the contents of a symbolic link.  Symbolic links may span
     file systems and may refer to directories.3、  Given one or two arguments, ln creates a link to an existing file source_file.  If link_name is given, the link has that name;
     link_name may also be a directory in which to place the link; otherwise it is placed in the current directory.  If only the
     directory is specified, the link will be made to the last component of source_file.
     Given more than two arguments, ln makes links in link_dirname to all the named source files.  The links made will have the same
     name as the files being linked to.
     When the utility is called as link, exactly two arguments must be supplied, neither of which may specify a directory.  No
     options may be supplied in this simple mode of operation, which performs a link(2) operation using the two passed arguments.