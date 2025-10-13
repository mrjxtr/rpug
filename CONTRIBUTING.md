# Contributing to RPUG

We're always looking to make this project better — whether it's adding more authentic Filipino names, fixing bugs, or improving features. Every contribution helps!

## 🏗️ Project Structure

```bash
rpug/
├── bin/                    # Compiled binaries go here
├── data/
│   ├── data.json          # The good stuff (names, locations)
│   └── examples/          # Sample responses
├── internal/
│   ├── config/            # Config and env handling
│   ├── generator/         # The brain - generates users
│   └── server/            # HTTP server and routes
├── main.go                # Entry point
├── Makefile               # Build commands
└── README.md              # Project documentation
```

## 🤓 Tech Stack

- **Go 1.21+** - Because we like things fast and compiled
- **Chi Router** - Lightweight, idiomatic HTTP routing
- **godotenv** - For that sweet environment config
- **Pure Go** - No external dependencies for core logic

## 🛠️ Development Setup

```bash
# Clone your fork
git clone https://github.com/YOUR_USERNAME/rpug.git
cd rpug

# Create a branch for your changes
git checkout -b feature/your-feature-name

# Run the server
make run

# Run tests
make test
```

## 🛠️ Build Commands

```bash
# Run the server
make run

# Build for your machine
make build

# Build for Linux AMD64
make build_amd64

# Run tests
make test

# Clean build artifacts
make clean

# Update dependencies
make tidy

# Build everything fresh
make all
```

## 🎨 Adding More Filipino Names

The names and locations live in `data/data.json`. This is where you can help the most!

### Adding Names

When adding names, keep these in mind:

- Use authentic Filipino names that are commonly used
- Include a good mix of traditional and modern names
- Avoid controversial or inappropriate names
- Make sure names are properly capitalized

Example structure:

```json
{
  "names": {
    "titles": {
      "male": ["Mr"],
      "female": ["Ms", "Mrs"]
    },
    "male_first_names": [
      "Juan",
      "Jose",
      "Carlo"
    ],
    "female_first_names": [
      "Maria",
      "Ana",
      "Kyla"
    ],
    "last_names": [
      "Santos",
      "Reyes",
      "Cruz"
    ]
  }
}
```

### Adding Locations

When adding Philippine locations:

- Use proper region and city names
- Include cities from different parts of the Philippines (Luzon, Visayas, Mindanao)
- Verify spelling and accuracy

Example structure:

```json
{
  "locations": [
    {
      "region": "National Capital Region",
      "cities": ["Makati", "Manila", "Pasay", "Quezon"]
    }
  ]
}
```

## 🧪 Testing

Always run tests before submitting:

```bash
make test
```

If you're adding new features, it will much appreciated if you include tests!

## 📝 Code Style

Follow these guidelines:

- Write clean, readable Go code
- Add docstrings to all functions
- Keep functions focused and small
- Follow Go conventions and idioms
- Run `go fmt` before committing

## 🔄 Pull Request Process

1. **Fork the repo** and create your branch from `dev`
2. **Make your changes** with clear, focused commits
3. **Run tests** to make sure nothing breaks
4. **Update documentation** if needed
5. **Submit a PR** with a clear description of what you changed and why

### Commit Message Format

Keep it simple and use these prefixes:

- `feat:` - new features
- `fix:` - bug fixes
- `update:` - general improvements
- `refactor:` - code restructuring
- `docs:` - documentation changes
- `chore:` - maintenance tasks

Examples:

```bash
feat: add more Visayan cities to locations
fix: handle empty name list gracefully
docs: update installation instructions
```

## 🐛 Found a Bug?

Open an issue with:

- Clear description of the bug
- Steps to reproduce
- Expected vs actual behavior
- Your Go version and OS

## 💡 Feature Requests

Got ideas? Open an issue and let's discuss! We love hearing new ideas for making RPUG better.

## 📄 License

By contributing, you agree that your contributions will be licensed under the MIT License.

---

### **Thanks for contributing! 🙏**

Every PR, issue, and suggestion helps make RPUG better for everyone. Let's build something awesome together! 💪
