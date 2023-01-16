### Chapter 5 - File I/O: Further details

System calls are atomic by design

## 5.1

fopen flags:
    O_CREAT + O_EXCL together assure atomicity on checking existence and creation
    

write flags:
    O_APPEND assures append atomicity
    lseek file distance is local information. does not reflect concurrent write
    if we do 2 steps: lseek to EOF + write, we can have a bad interleaving from 2 processes
    p1 = lseek
    switch
    p2 = lseek (correct)
    p2 = write (new lseek)
    switch
    p1 = write (from wrong lseek)
    BANG

    solution: use write with O_APPEND option =)

## 5.2

fctl = file control
many possibilities, will analyse throughout book

### 5.3

fctl usages: get access modes, file flags, modify file flags

### 5.4

File descriptors have 2 different contexts: per process, and system wide
There is also the inodes

* per process file descriptor table => multiple FDs can point to same file
* system wide table of open FDs => multiple process FDs can point to a single entry here. One entry here maps 1:n to an inode (multiple table entries to same inode)
* filesystem i-node tables => maps to actual files

offsets belong to a single table entry, so multiple process FDs can share the share the same offset. beware

## 5.5

close(newfd); newfd=dup(oldfd); => dup will duplicate fd oldfd to newfd

newFd = dup2(oldfd,newfd) => closes newfd if it is open and clones oldfd on newfd. the closing fails silently

## 5.6

pread => atomically reads from a specific offset and returns file offset to the original position

bytesread= pread(fd, buf, count, offset)
byteswritten = pwrite(fd, buf, count, offset)

pread is equivalent to running the following atomically:

off_t orig;
orig = lseek(fd, 0, SEEK_CUR);
lseek(fd,offset, SEEK_SET);
s = read(fd,buf,len);
lseek(fd,orig,SEEK_SET);

useful for multithreaded scenario on same process, or multiple process pointing to same file table entry =)

## 5.7

readv and writev => scatter/gather io, atomically

ssize_t preadv(int fd, const struct iovec *iov, int iovcnt, off_t offset);
ssize_t preadv(int fd, const struct iovec *iov, int iovcnt, off_t offset);

## 5.8

sets size of a file to a given length. difference is in how file is referenced

int truncate(const char *pathname, off_t length);
int ftruncate(int fd, off_t length);

file longer than length > data is lost
file shorter than length > filled with holes or null bytes

## 5.9

O_NONBLOCK flag for opening files:
* If the file can’t be opened immediately, then open() returns an error instead of
blocking. One case where open() can block is with FIFOs (Section 44.7).
* After a successful open(), subsequent I/O operations are also nonblocking. If
an I/O system call can’t complete immediately, then either a partial data transfer is performed or the system call fails with one of the errors EAGAIN or
EWOULDBLOCK. Which error is returned depends on the system call. On Linux, as
on many UNIX implementations, these two error constants are synonymous.

Nonblocking mode can be used with devices (e.g., terminals and pseudoterminals),
pipes, FIFOs, and sockets. (Because file descriptors for pipes and sockets are not
obtained using open(), we m


## 5.10

files larger than 2GB will overflow the offset value in 32 bit systems. beware (not gonna bother cuz nobody runs this shit anymore)

## 5.11

/dev/fd will hold all fds for the process that calls it
it is a symbolic link to /proc/$(procid)

## 5.12

returns FD for a temp file
template must be a string that ends with XXXXXX => will generate random chars and assure uniqueness
tempate must be char* because will be modified, not constant string


int mkstemp(char *template);

The mkstemp() function creates the file with read and write permissions for the
file owner (and no permissions for other users), and opens it with the O_EXCL flag,
guaranteeing that the caller has exclusive access to the file.
Typically, a temporary file is unlinked (deleted) soon after it is opened, using
the unlink() system call (Section 18.3). 

the name will vanish from filesystem but the file itself will only be deleted when no more FDs point to it


