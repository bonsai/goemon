import argparse
import sys
import os
import json
from i2t_watcher import VlmModelAgent

def main():
    parser = argparse.ArgumentParser(description="Watcher CLI for Go integration")
    parser.add_argument("--image", type=str, required=True, help="Path to the image to analyze")
    parser.add_argument("--prompt", type=str, default="Describe this image in detail.", help="VLM prompt")
    parser.add_argument("--model_path", type=str, default="C:/models/vlm/moondream2", help="Local model path")
    args = parser.parse_args()

    if not os.path.exists(args.image):
        print(f"Error: Image not found: {args.image}", file=sys.stderr)
        sys.exit(1)

    try:
        # VLM インスタンス化
        agent = VlmModelAgent(model_id=args.model_path)
        
        # 解析実行
        caption = agent.analyze_image(args.image, args.prompt)
        
        # 結果をJSON形式で出力
        result = {
            "caption": caption,
            "model": "moondream2",
            "image_path": args.image
        }
        print(json.dumps(result))
        
    except Exception as e:
        print(f"Error: {e}", file=sys.stderr)
        sys.exit(1)

if __name__ == "__main__":
    main()
