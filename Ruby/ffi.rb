require 'ffi'

module Lib
  extend FFI::Library
  ffi_lib './C/turbosizenator.so'
  attach_function :ConvertImageFromFilePath, [:string, :string, :string], :void
end

Lib.ConvertImageFromFilePath("./tests/image.png", "image.gif", "gif")
