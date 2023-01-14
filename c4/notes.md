### Chapter 4 - File I/O: The universal I/O model

Basic FDs: 0=stdin, 1=stdout, 2=stderr

Basic syscalls:

    fd = open(pathname, flags, mode) => opens file and returns FD (success) or -1 (error)
    numread = read(fd,buffer,count) => reads max of count bytes into a buffer. returns number of bytes read or -1 (fail)
    status = close(fd) => releases the FD and closes the file
    new_offset = lseek(fd,offset,whence) => changes the base pointer of the fd by offset from start of file/current offset/end of file depending on whence flag
    ioctl => ops outside of standard io model, to be revised later
