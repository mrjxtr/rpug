# RANDOM PINOY USER GENERATOR (RPUG)

> _"Pinoy test data ba kailangan mo?! Heto na!"_ 🇵🇭
> ℹ️ The API is live and ready to use at **[https://randompinoy.xyz/api/v1/pinoys](https://randompinoy.xyz/api/v1/pinoys)** — no installation required! 🚀

Generate realistic Filipino user data faster than you can say "Mabuhay!" Perfect for testing, demos, or when you need fake Pinoy users that actually look and feel legit.

**100% Free. 100% Open Source. 100% Pinoy.** 🆓

## 🎯 What's This All About?

RPUG is a lightweight REST API that generates random Filipino user profiles. No more boring "John Doe" or "Jane Smith" in your test data — get authentic Pinoy names, real Philippine locations, and data that actually makes sense for Filipino users.

Built with Go, powered by Filipino spirit. ✨

## ✨ Features

- **Free & Open Source** - Use it anywhere, anytime. No API keys, no BS
- **Authentic Filipino Names** - From Juan dela Cruz to Princess Mae Villanueva
- **Real Philippine Locations** - Cities and regions from Luzon to Mindanao
- **Deterministic Seeds** - Same seed = same data (perfect for reproducible tests)
- **Flexible Results** - Generate 1 to 1,000 users in a single request
- **Fast & Lightweight** - Because ain't nobody got time for slow APIs
- **JSON All The Way** - Easy to parse, easy to use

## 🌐 Live API Usage

🔗 **[https://randompinoy.xyz/api/v1/pinoys](https://randompinoy.xyz/v1/pinoys)**

### Try It Now

```bash
# Generate 1 user
curl https://randompinoy.xyz/api/v1/pinoys

# Generate 10 users
curl https://randompinoy.xyz/api/v1/pinoys?results=10

# Use a seed for reproducible data
curl https://randompinoy.xyz/api/v1/pinoys?seed=2d0cd4170d54fbacdcc1e679ecf394cd
```

### Want to Run It Locally?

Check out the [Quick Start Guide](CONTRIBUTING.md#🛠️-development-setup) in our contributing docs if you want to run your own instance or contribute to the project.

## 📡 API Endpoints

### Health Check

```bash
GET /ping
```

Returns `200 OK` if the server is alive and kicking.

### Generate Random Users

```bash
GET /api/v1/pinoys
```

Generate random Filipino user profiles. That's it. That's the API.

## 🎮 Usage Examples

### Basic Request (1 user)

```bash
curl https://randompinoy.xyz/api/v1/pinoys
```

### Generate Multiple Users

```bash
# Get 5 users
curl https://randompinoy.xyz/api/v1/pinoys?results=5

# Go crazy with 1000 users
curl https://randompinoy.xyz/api/v1/pinoys?results=1000
```

### Use a Seed for Reproducible Data

```bash
# Same seed = same data every time
curl https://randompinoy.xyz/api/v1/pinoys?seed=2d0cd4170d54fbacdcc1e679ecf394cd
```

> **Note:** If you're running locally, replace `https://randompinoy.xyz` with `http://localhost:3000`

## 📦 Response Format

> As of 2025-10-14, some data will be blank and are still being implemented

```json
{
  "results": [
    {
      "name": {
        "title": "Mr",
        "first": "Carlo",
        "last": "Santos"
      },
      "dob": {
        "date": "1989-05-30T23:07:31.851Z",
        "age": 36
      },
      "location": {
        "street": {
          "number": 5843,
          "name": "Duke St"
        },
        "city": "Pagadian",
        "region": "Zamboanga Del Sur",
        "country": "Philippines",
        "zipcode": "7016"
      },
      "gender": "male",
      "phone": "09123456789",
      "email": "carlo.santos@example.com",
      "login": {
        "uuid": "14fa0589-a264-4fdb-945b-4971d138f118",
        "username": "carlo.santos123",
        "password": "secret"
      },
      "registered": {
        "date": "2025-04-03T02:01:00.708Z",
        "age": 3
      }
    }
  ],
  "info": {
    "seed": "2d0cd4170d54fbacdcc1e679ecf394cd",
    "results": 1,
    "version": "0.1.x-alpha"
  }
}
```

## 🔧 Query Parameters

| Parameter | Type   | Default | Max  | Description                    |
| --------- | ------ | ------- | ---- | ------------------------------ |
| `results` | int    | 1       | 1000 | Number of users to generate    |
| `seed`    | string | random  | -    | Seed for deterministic results |

**Pro tip:** Results are clamped between 1-1000. Use the `results` parameter to get multiple users in one request instead of making rapid-fire requests.

## 🚦 Rate Limiting

To keep the API fast and fair for everyone, we enforce these limits:

- **60 requests per minute** per IP address (~1 request per second average)

If you hit the limit, you'll get a `429 Too Many Requests` response. Just wait a moment and try again, or better yet — use the `results` parameter to get multiple users in a single request!

## 📝 Notes

- Email generation is not yet implemented (placeholder for now)
- Phone number format coming soon
- Date of birth and registration dates are placeholders
- Pagination support is on the roadmap

This is a work in progress, pero functional na siya! Ship it! 🚢

## 📄 License

This project is **free and open source** under the GNU General Public License v3.0 (GPL-3.0). Use it, fork it, share it — walang bayad! Check the `LICENSE` file for the full details.

Just give credit where it's due, okay? 😉

## 🤝 Contributing

Want to add more Filipino names? Found a bug? Got ideas for features?

Check out [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines on how to contribute. PRs and issues are welcome! 🙏

---

**Made with ❤️ and kape ☕ by [@mrjxtr](https://mrjxtr.dev)**
**Mabuhay Pinoy developers! 🇵🇭**
