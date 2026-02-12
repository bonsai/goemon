import sys
import json

def generate_spell(seed, fragments):
    """
    GLM (Generative Language Model) を使用して、
    シードとDBから取得した断片を元に新しい『呪文（SPELL）』を生成する。
    """
    # 実際の実装ではここでローカルのLLM APIを叩くなどの処理を行う
    # 現在はモックとして生成結果を返す
    
    prompt = f"Seed: {seed}\nFragments: {fragments}"
    
    # 仮の生成ロジック
    spell = f"Casted Spell from '{seed}' using magic fragments: {', '.join(fragments)}"
    
    result = {
        "status": "success",
        "spell": spell,
        "raw_prompt": prompt
    }
    return result

if __name__ == "__main__":
    if len(sys.argv) < 3:
        print(json.dumps({"status": "error", "message": "Usage: python glm_bridge.py <seed> <fragments_json>"}))
        sys.exit(1)
        
    seed = sys.argv[1]
    try:
        fragments = json.loads(sys.argv[2])
    except:
        fragments = [sys.argv[2]] # 文字列として扱うフォールバック
        
    response = generate_spell(seed, fragments)
    print(json.dumps(response))
