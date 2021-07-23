## import packages, if they don't have em it'll tell them to get them.
import sys
try:
    import PIL.Image
except:
    print("Pillow is not installed! Install it using the following command: 'pip3 install pillow'")
    quit()

#   Chars to use with making the image, from most intense to least intense
ASCII_CHARS = ["@","#","S","%","?","*","+",";",":",",","."," "]

#   Resize the image to make it easier to work with
def image_resize(image,new_width=50):
    width, height = image.size
    ratio = height/width
    new_height = int(new_width * ratio)
    resized_image = image.resize((new_width,new_height))
    return(resized_image)

#   Grayscale the image
def convert_to_grayscale(image):
    grayscale_image = image.convert("L")
    return(grayscale_image)

#   Converts the pixels to the characters on the table above. Returns the string of pixel to characters.
def convert_to_ascii(image):
    pixels = image.getdata()
    characters = "".join([ASCII_CHARS[pixel//25] for pixel in pixels])
    return(characters)


def main(image=PIL.Image.open(sys.argv[1]),new_width=50):
    
    new_image_in_chars = convert_to_ascii(convert_to_grayscale(image_resize(image)))
    
    len_of_chars = len(new_image_in_chars)
    final_product = "\n".join(new_image_in_chars[i:(i+new_width)] for i in range(0,len_of_chars,new_width))
    
    print(final_product)

    with open("ascii-image.txt", "w") as doc:
        doc.write(final_product)
        print("Finished!")

main()