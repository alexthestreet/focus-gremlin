# Focus Gremlin 👹

A tiny, persistent desktop companion whose only job is to ask:

> “Are you actually working right now?”

Focus Gremlin is a lightweight accountability tool that runs on your machine and periodically checks in to make sure you haven't drifted into the warm embrace of distractions, procrastination, or existential wandering.

It does not manage your life.
It does not make decisions for you.
It just… watches. And asks questions.

Relentlessly.

---

## ✨ Features

* ⏰ **Hourly Check-ins**
  Pops up at a configurable interval (default: every hour) to ask if you're on track.

* 🧠 **Self-Accountability Prompts**
  Forces you to acknowledge your current state:

  * On track ✅
  * Off track ⚠️
  * Deep in the void 🌪️

* 📝 **Simple Logging (Optional)**
  Records your responses locally so you can review how your day actually went (brace yourself).

* 🔕 **Easy On/Off**
  Run it when you want accountability. Close it when you want plausible deniability.

* ⚡ **Lightweight**
  Built in Go. Fast, minimal, no bloated productivity suite pretending to change your life.

---

## 🧩 Philosophy

Focus Gremlin exists because:

* You already know what you need to do
* You don’t need another complex system
* You just need something to interrupt the drift

This is not a planner.
This is not a task manager.
This is a **mirror with a timer**.

---

## 🚀 How It Works

1. Start the app
2. It runs quietly in the background
3. Every interval, it interrupts you with a check-in prompt
4. You respond honestly (or lie to yourself, your call)
5. Repeat until productivity or guilt wins

---

## 🛠️ Setup

```bash
# clone the repo
git clone https://github.com/alexthestreet/focus-gremlin.git

cd focus-gremlin

# build
go build -o focus-gremlin

# run
./focus-gremlin
```

---

## ⚙️ Configuration (Optional)

You can tweak behavior via flags or config (depending on how ambitious you get):

* `--interval` → check-in frequency (e.g., 30m, 1h)
* `--start` / `--end` → active hours
* `--log` → enable/disable logging
* `--silent` → reduce notifications (for when you're fragile)

---

## 🧪 Future Ideas

If you decide to spiral deeper into this:

* Integration with ChatGPT for “I’m stuck” help
* Desktop notifications vs. modal popups
* Stats dashboard (a.k.a. your accountability report card)
* Passive-aggressive message modes
* “Gremlin gets angrier over time” difficulty scaling

---

## ⚠️ Disclaimer

Focus Gremlin:

* Will not fix your motivation
* Will not stop you from ignoring it
* Will not physically prevent you from opening YouTube

It will, however, make you *aware* that you're doing those things.

And honestly, that’s already more than most tools.
