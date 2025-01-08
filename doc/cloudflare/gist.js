const html = `<!DOCTYPE html>
<html lang="zh">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>BESTRUI</title>
    <style>
        :root {
            --primary-gradient: linear-gradient(45deg, #12c2e9, #c471ed, #f64f59);
            --glass-bg: rgba(255, 255, 255, 0.1);
            --glass-border: rgba(255, 255, 255, 0.18);
            --glass-shadow: rgba(31, 38, 135, 0.37);
        }

        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }

        body { 
            min-height: 100vh;
            display: flex;
            justify-content: center;
            align-items: center;
            background: var(--primary-gradient);
            background-size: 400% 400%;
            animation: gradient 15s ease infinite;
            font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, sans-serif;
            -webkit-font-smoothing: antialiased;
        }

        @keyframes gradient {
            0% { background-position: 0% 50%; }
            50% { background-position: 100% 50%; }
            100% { background-position: 0% 50%; }
        }

        .container {
            text-align: center;
            color: white;
            padding: 2.5rem;
            border-radius: 2rem;
            background: var(--glass-bg);
            backdrop-filter: blur(10px);
            -webkit-backdrop-filter: blur(10px);
            box-shadow: 0 8px 32px 0 var(--glass-shadow);
            border: 1px solid var(--glass-border);
            max-width: 90vw;
            width: 340px;
            transition: transform 0.3s ease;
        }

        .container:hover {
            transform: translateY(-5px);
        }

        h1 {
            font-size: clamp(2rem, 8vw, 2.5rem);
            margin: 0 0 1rem 0;
            font-weight: 500;
            letter-spacing: 3px;
            text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.2);
        }

        .logo {
            width: 70px;
            height: 70px;
            margin: 0 auto 1.5rem;
            position: relative;
        }

        .circle {
            position: absolute;
            width: 100%;
            height: 100%;
            border-radius: 50%;
            border: 3px solid white;
            animation: rotate 8s linear infinite;
            opacity: 0.8;
        }

        .circle:nth-child(1) { 
            animation-delay: -2s; 
            border-color: rgba(255, 255, 255, 0.9);
        }
        .circle:nth-child(2) { 
            animation-delay: -4s;
            border-color: rgba(255, 255, 255, 0.7);
        }
        .circle:nth-child(3) { 
            animation-delay: -6s;
            border-color: rgba(255, 255, 255, 0.5);
        }

        @keyframes rotate {
            0% { transform: rotate(0deg) scale(0.8); }
            50% { transform: rotate(180deg) scale(1.2); }
            100% { transform: rotate(360deg) scale(0.8); }
        }

        .quote {
            font-size: 1.1rem;
            opacity: 0.9;
            margin: 1.2rem 0;
            font-style: italic;
            font-weight: 300;
            text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.1);
        }

        @media (max-width: 480px) {
            .container {
                padding: 2rem;
                width: 300px;
            }
            .logo {
                width: 60px;
                height: 60px;
            }
            .quote {
                font-size: 1rem;
            }
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="logo">
            <div class="circle"></div>
            <div class="circle"></div>
            <div class="circle"></div>
        </div>
        <h1>BESTRUI</h1>
        <div class="quote">Exploring the digital frontier</div>
    </div>
</body>
</html>`;

// 验证 token 的函数
async function validateToken(url, env) {
    const token = url.searchParams.get('token');
    if (!token) {
        return false;
    }
    return token === env.AUTH_TOKEN;
}

// 处理请求的主函数
export default {
    async fetch(request, env) {
        try {
            const url = new URL(request.url);

            // 处理根路径请求，返回 HTML 页面
            if (url.pathname === '/') {
                return new Response(html, {
                    headers: { 'Content-Type': 'text/html' },
                });
            }

            // GitHub API 反向代理
            if (url.pathname.startsWith('/github/')) {
                // 重构 GitHub API 请求路径
                const githubPath = url.pathname.replace('/github/', '');
                const githubUrl = `https://api.github.com/${githubPath}`;

                // 转发原始请求头
                const headers = new Headers(request.headers);
                headers.set('User-Agent', 'Cloudflare-Worker');

                // 转发请求到 GitHub API
                const githubResponse = await fetch(githubUrl, {
                    method: request.method,
                    headers: headers,
                    body: request.method !== 'GET' ? await request.text() : undefined
                });

                // 构建响应头
                const responseHeaders = new Headers({
                    'Access-Control-Allow-Origin': '*',
                    'Access-Control-Allow-Methods': 'GET, POST, PUT, DELETE, OPTIONS',
                    'Access-Control-Allow-Headers': 'Content-Type, Authorization',
                    'Content-Type': 'application/json'
                });

                // 转发 GitHub API 的响应
                return new Response(await githubResponse.text(), {
                    status: githubResponse.status,
                    headers: responseHeaders
                });
            }

            if (url.pathname === '/bestrui') {
                // 验证 token
                if (!await validateToken(url, env)) {
                    return new Response(JSON.stringify({
                        code: 401,
                        message: '未授权访问'
                    }), {
                        status: 401,
                        headers: { 'Content-Type': 'application/json' }
                    });
                }
                const key = url.searchParams.get('key');
                //获取时间戳
                const timestamp = Date.now();
                const gistContent = await fetch(`https://gist.githubusercontent.com/${env.GITHUB_USER}/${env.GITHUB_ID}/raw/${key}?timestamp=${timestamp}`).then(res => res.text());
                return new Response(gistContent, {
                    headers: { 'Content-Type': 'text/plain; charset=utf-8' }
                });
            }

            return new Response(JSON.stringify({
                code: 404,
                message: '404 Not Found'
            }), {
                status: 404,
                headers: { 'Content-Type': 'application/json' }
            });
        } catch (error) {
            return new Response('发生错误: ' + error.message, { status: 500 });
        }
    }
};
