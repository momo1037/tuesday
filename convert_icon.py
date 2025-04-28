from PIL import Image
import os
import sys

def convert_png_to_ico(input_path, output_path):
    """
    将PNG图片转换为ICO格式
    :param input_path: 输入PNG文件路径
    :param output_path: 输出ICO文件路径
    """
    try:
        img = Image.open(input_path)
        img.save(output_path, format='ICO', sizes=[(32,32), (48,48), (64,64), (128,128)])
        print(f"成功将 {input_path} 转换为 {output_path}")
        return True
    except Exception as e:
        print(f"转换失败: {e}")
        return False

if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("用法: python convert_icon.py <input.png> [output.ico]")
        sys.exit(1)
    
    input_file = sys.argv[1]
    output_file = sys.argv[2] if len(sys.argv) > 2 else "icon.ico"
    
    if not os.path.exists(input_file):
        print(f"错误: 文件 {input_file} 不存在")
        sys.exit(1)
    
    if convert_png_to_ico(input_file, output_file):
        print("转换完成，可以使用以下命令将ICO嵌入到Go程序中:")
        print(f"go generate")
    else:
        sys.exit(1)