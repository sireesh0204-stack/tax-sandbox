package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>ProConnect Tax Engine</title>
    <style>
        * {
            box-sizing: border-box;
            margin: 0;
            padding: 0;
        }
        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
            background: #0f172a;
            color: #f1f5f9;
            line-height: 1.6;
            min-height: 100vh;
        }
        .container {
            display: flex;
            min-height: 100vh;
        }
        .sidebar {
            width: 280px;
            background: rgba(30, 41, 59, 0.95);
            backdrop-filter: blur(20px);
            border-right: 1px solid rgba(255, 255, 255, 0.1);
            padding: 2rem 1.5rem;
            display: flex;
            flex-direction: column;
            gap: 0.5rem;
        }
        .logo {
            font-size: 1.25rem;
            font-weight: 600;
            color: #38bdf8;
            margin-bottom: 2rem;
            letter-spacing: -0.02em;
        }
        .nav-item {
            padding: 0.75rem 1rem;
            border-radius: 12px;
            cursor: pointer;
            transition: all 0.2s cubic-bezier(0.23, 1, 0.32, 1);
            font-size: 0.9rem;
            display: flex;
            align-items: center;
            gap: 0.75rem;
            color: #94a3b8;
        }
        .nav-item:hover {
            background: rgba(56, 189, 248, 0.1);
            color: #38bdf8;
            transform: translateX(4px);
        }
        .nav-item.active {
            background: rgba(56, 189, 248, 0.15);
            color: #38bdf8;
            border-left: 3px solid #38bdf8;
        }
        .main-content {
            flex: 1;
            padding: 2rem 3rem;
            overflow-y: auto;
        }
        .form-section {
            background: rgba(30, 41, 59, 0.6);
            border: 1px solid rgba(255, 255, 255, 0.08);
            border-radius: 16px;
            padding: 2rem;
            margin-bottom: 1.5rem;
            transition: all 0.3s cubic-bezier(0.23, 1, 0.32, 1);
        }
        .form-section:hover {
            border-color: rgba(56, 189, 248, 0.3);
            transform: translateY(-1px);
        }
        .form-section-title {
            font-size: 1.25rem;
            font-weight: 600;
            margin-bottom: 1.5rem;
            color: #38bdf8;
        }
        .form-grid {
            display: grid;
            grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
            gap: 1.5rem;
        }
        .form-group {
            display: flex;
            flex-direction: column;
            gap: 0.5rem;
        }
        .form-label {
            font-size: 0.875rem;
            font-weight: 500;
            color: #cbd5e1;
        }
        .form-input {
            background: rgba(15, 23, 42, 0.6);
            border: 1px solid rgba(255, 255, 255, 0.1);
            border-radius: 10px;
            padding: 0.75rem 1rem;
            color: #f1f5f9;
            font-size: 0.95rem;
            transition: all 0.2s cubic-bezier(0.23, 1, 0.32, 1);
            outline: none;
        }
        .form-input:focus {
            border-color: #38bdf8;
            transform: translateY(-1px);
            box-shadow: 0 0 0 3px rgba(56, 189, 248, 0.2);
        }
        .btn {
            padding: 0.75rem 1.5rem;
            border-radius: 10px;
            font-weight: 600;
            font-size: 0.9rem;
            cursor: pointer;
            transition: all 0.2s cubic-bezier(0.23, 1, 0.32, 1);
            border: none;
            outline: none;
        }
        .btn-primary {
            background: #38bdf8;
            color: #0f172a;
        }
        .btn-primary:hover {
            background: #0ea5e9;
            transform: translateY(-1px);
        }
        .btn-secondary {
            background: rgba(255, 255, 255, 0.1);
            color: #f1f5f9;
        }
        .btn-secondary:hover {
            background: rgba(255, 255, 255, 0.15);
        }
        @media (max-width: 1024px) {
            .container {
                flex-direction: column;
            }
            .sidebar {
                width: 100%;
            }
        }
        @media (prefers-reduced-motion: reduce) {
            *, *::before, *::after {
                animation-duration: 0.01ms !important;
                animation-iteration-count: 1 !important;
                transition-duration: 0.01ms !important;
            }
        }
    </style>
</head>
<body>
    <div class="container">
        <aside class="sidebar">
            <div class="logo">ProConnect Tax · 2026</div>
            <div style="margin-bottom: 1.5rem;">
                <div class="nav-label" style="margin-bottom: 0.5rem;">Income Returns</div>
                <div class="nav-item" onclick="selectSection(this, 'overview')">📊 Overview</div>
                <div class="nav-item" onclick="selectSection(this, 'personal')">👤 Personal</div>
                <div class="nav-item" onclick="selectSection(this, 'income')">💰 Income</div>
            </div>
            <div style="margin-bottom: 1.5rem;">
                <div class="nav-label" style="margin-bottom: 0.5rem;">Credits & Taxes</div>
                <div class="nav-item" onclick="selectSection(this, 'deductions')">📎 Credits</div>
                <div class="nav-item" onclick="selectSection(this, 'taxes')">💸 Taxes</div>
            </div>
            <div>
                <div class="nav-label" style="margin-bottom: 0.5rem;">Actions</div>
                <div class="nav-item" onclick="selectSection(this, 'mef')">📤 Submit</div>
            </div>
        </aside>
        <main class="main-content">
            <h1 style="font-size: 2rem; margin-bottom: 0.5rem; letter-spacing: -0.02em;">Form 1040: High Net Worth 2026</h1>
            <p style="color: #94a3b8;">David Vance · Cohort 6+ Years Tax Technology</p>
            
            <div id="content">
                <div class="form-section">
                    <h2 class="form-section-title">Return Overview</h2>
                    <div class="form-grid">
                        <div class="form-group">
                            <label class="form-label">Filer Name</label>
                            <div class="form-input">David Vance</div>
                        </div>
                        <div class="form-group">
                            <label class="form-label">Total Income</label>
                            <div class="form-input" style="color: #38bdf8; font-weight: 600;">$1,250,000+</div>
                        </div>
                        <div class="form-group">
                            <label class="form-label">Net Income</label>
                            <div class="form-input">$980,000+</div>
                        </div>
                        <div class="form-group">
                            <label class="form-label">Estimated Tax Liability</label>
                            <div class="form-input">$320,000+</div>
                        </div>
                    </div>
                </div>
            </div>
        </main>
    </div>
    <script>
        const sections = {
            overview: `
                <div class="form-section">
                    <h2 class="form-section-title">Return Overview</h2>
                    <div class="form-grid">
                        <div class="form-group">
                            <label class="form-label">Filer Name</label>
                            <div class="form-input">David Vance</div>
                        </div>
                        
