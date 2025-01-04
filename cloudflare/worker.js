// 定义 HTML 模板
const html = `<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>BESTRUI</title>
    <style>
        body { 
            margin: 0;
            min-height: 100vh;
            display: flex;
            justify-content: center;
            align-items: center;
            background: linear-gradient(45deg, #12c2e9, #c471ed, #f64f59);
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
            padding: 2rem;
            border-radius: 1.5rem;
            background: rgba(255, 255, 255, 0.1);
            backdrop-filter: blur(10px);
            -webkit-backdrop-filter: blur(10px);
            box-shadow: 0 8px 32px 0 rgba(31, 38, 135, 0.37);
            border: 1px solid rgba(255, 255, 255, 0.18);
            max-width: 90vw;
            width: 300px;
        }

        h1 {
            font-size: min(2.5rem, 10vw);
            margin: 0 0 1rem 0;
            font-weight: 300;
            letter-spacing: 2px;
        }

        .logo {
            width: 60px;
            height: 60px;
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
        }

        .circle:nth-child(1) { animation-delay: -2s; }
        .circle:nth-child(2) { animation-delay: -4s; }
        .circle:nth-child(3) { animation-delay: -6s; }

        @keyframes rotate {
            0% { transform: rotate(0deg) scale(0.8); }
            50% { transform: rotate(180deg) scale(1.2); }
            100% { transform: rotate(360deg) scale(0.8); }
        }

        .quote {
            font-size: 1rem;
            opacity: 0.8;
            margin: 1rem 0;
            font-style: italic;
        }

        @media (max-width: 480px) {
            .container {
                padding: 1.5rem;
            }
            .logo {
                width: 50px;
                height: 50px;
            }
            .quote {
                font-size: 0.9rem;
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

            if (url.pathname === '/storage') {
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

                // GET 请求用于读取数据
                if (request.method === 'GET') {
                    const key = url.searchParams.get('key');
                    if (!key) {
                        return new Response(JSON.stringify({
                            code: 400,
                            message: '请提供键名'
                        }), {
                            status: 400,
                            headers: { 'Content-Type': 'application/json' }
                        });
                    }

                    try {
                        const object = await env.SUB_BUCKET.get(key);

                        if (object === null) {
                            return new Response(JSON.stringify({
                                code: 404,
                                message: '未找到该键对应的值'
                            }), {
                                status: 404,
                                headers: { 'Content-Type': 'application/json' }
                            });
                        }

                        const data = await object.text();
                        return new Response(data, {
                            headers: { 'Content-Type': 'text/plain; charset=utf-8' }
                        });
                    } catch (error) {
                        return new Response(JSON.stringify({
                            code: 500,
                            message: '读取数据失败',
                            error: error.message
                        }), {
                            status: 500,
                            headers: { 'Content-Type': 'application/json' }
                        });
                    }
                }

                // POST 请求用于写入数据
                if (request.method === 'POST') {
                    const { key, value } = await request.json();
                    if (!key || !value) {
                        return new Response(JSON.stringify({
                            code: 400,
                            message: '请提供键和值'
                        }), {
                            status: 400,
                            headers: { 'Content-Type': 'application/json' }
                        });
                    }

                    try {
                        await env.SUB_BUCKET.put(key, value);
                        return new Response(JSON.stringify({
                            code: 200,
                            message: '数据写入成功'
                        }), {
                            status: 200,
                            headers: { 'Content-Type': 'application/json' }
                        });
                    } catch (error) {
                        return new Response(JSON.stringify({
                            code: 500,
                            message: '数据写入失败',
                            error: error.message
                        }), {
                            status: 500,
                            headers: { 'Content-Type': 'application/json' }
                        });
                    }
                }
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
