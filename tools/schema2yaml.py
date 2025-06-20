import os

schemas_path = "/home/joseph/coding_base/drawn-by-fate/pkg/config/package/configs"
for root, _, files in os.walk(schemas_path):
    for fname in files:
        if fname.endswith(".json"):
            path = os.path.join(root, fname)
            # os.system(f"j2y {path}")
            os.remove(path)
