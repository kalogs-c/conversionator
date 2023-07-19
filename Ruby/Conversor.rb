require 'ffi'

module Conversor
  extend FFI::Library
  ffi_lib './C/turbosizenator.so'
  attach_function :ConvertImageFromFilePath, [:string, :string, :string], :void
  attach_function :ConvertImageFromBytes, [:pointer, :int, :string], :pointer

  def self.FromPath(image_path, image_dest, target_format)
    ConvertImageFromFilePath(image_path, image_dest, format)
  end

  def self.FromBytes(file_data, target_format)
    fsize = file_data.length
    
    converted_image_ptr = ConvertImageFromBytes(file_data, fsize, target_format)

    converted_image_data = converted_image_ptr.read_array_of_type(:int, :read_uint8, fsize)

    converted_image_data
  end
end

original = File.open('./tests/image.png', 'rb')
bytes = original.read
image_bytes = Conversor.FromBytes(bytes, 'gif')
puts image_bytes

File.binwrite("image.gif", image_bytes)
