// The module 'vscode' contains the VS Code extensibility API
// Import the module and reference it with the alias vscode in your code below
import * as vscode from 'vscode';

// This method is called when your extension is activated
// Your extension is activated the very first time the command is executed
export function activate(context: vscode.ExtensionContext) {
	
	context.subscriptions.push(vscode.window.registerWebviewViewProvider(
		"devsync.main",
		{
			"resolveWebviewView":(webview) => {
				webview.webview.html = `<html>
					<body>
						<h1>hello</h1>
					</body>
				</html>`;
				const panel = vscode.window.createWebviewPanel("DevSync", "DevSync", vscode.ViewColumn.One, {});
				panel.webview.html = `
					<html>
						<body>
							<h1>hi</h1>
						</body>
					</html>
				`;
			}
		}
	))

}

// This method is called when your extension is deactivated
export function deactivate() {}
