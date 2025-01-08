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
                    const filename = url.searchParams.get('filename');
                    if (!filename) {
                        return new Response(JSON.stringify({
                            code: 400,
                            message: '请提供文件名'
                        }), {
                            status: 400,
                            headers: { 'Content-Type': 'application/json' }
                        });
                    }

                    try {
                        const object = await env.SUB_BUCKET.get(filename);

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
                    const { filename, value } = await request.json();
                    if (!filename || !value) {
                        return new Response(JSON.stringify({
                            code: 400,
                            message: '请提供文件名和值'
                        }), {
                            status: 400,
                            headers: { 'Content-Type': 'application/json' }
                        });
                    }

                    try {
                        await env.SUB_BUCKET.put(filename, value);
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
