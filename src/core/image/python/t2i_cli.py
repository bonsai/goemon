import argparse
import sys
import os
from t2i_baker import Baker

def main():
    parser = argparse.ArgumentParser(description="Baker CLI for Go integration")
    parser.add_argument("--prompt", type=str, required=True, help="Image generation prompt")
    parser.add_argument("--id", type=str, required=True, help="Image ID for tracking")
    args = parser.parse_args()

    # 出力先を ADR-0030 に合わせる
    output_dir = "data/images"
    os.makedirs(output_dir, exist_ok=True)
    
    # Baker インスタンス化 (Stable Diffusion)
    # 本来は GPU メモリ節約のためシングルトン的に管理すべきだが、
    # プロセス分離の場合はここでロードされる
    baker = Baker(model_id="runwayml/stable-diffusion-v1-5")
    
    # 画像生成
    # baker.bake 内で保存されるが、IDを指定してリネームする
    temp_path = baker.bake(args.prompt)
    final_path = os.path.join(output_dir, f"{args.id}.png")
    
    if os.path.exists(temp_path):
        os.rename(temp_path, final_path)
        print(final_path) # Go 側でキャッチする
    else:
        print(f"Error: Failed to generate image", file=sys.stderr)
        sys.exit(1)

if __name__ == "__main__":
    main()
