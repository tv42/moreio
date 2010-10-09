============================================================================
 ``moreio`` - combine a ReadCloser and a WriteCloser into a ReadWriteCloser
============================================================================

This is a simple utility package. Hopefully this will get integrated
into Go upstream ``io`` package.

If you have a ``ReadCloser`` and a ``WriteCloser``, you can combine
them (reads go to ``ReadCloser``, writes go to ``WriteCloser``, close
goes to both) with a simple call to ``NewReadWriteCloser``.

This is most useful for taking the read end of one pipe, the write end
of another, and making those appear like a single logical file.
