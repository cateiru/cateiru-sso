# フロントエンドテスト戦略

Oreore.meのフロントエンドは、Storybookを中心とした包括的なテスト戦略により品質を確保しています。

## 🧪 テスト戦略概要

### テストピラミッド

```mermaid
pyramid TB
    subgraph "E2E Tests"
        E2E[手動テスト<br/>Critical Path]
    end

    subgraph "Integration Tests"
        INT[Storybook Interactions<br/>Component Integration]
    end

    subgraph "Unit Tests"
        UNIT[Storybook Stories<br/>Component States]
        LOGIC[Utility Functions<br/>Business Logic]
    end
```

### テストツール構成

| レベル | ツール | 対象 | 実行タイミング |
|--------|--------|------|----------------|
| **Visual Testing** | Storybook | コンポーネント外観 | 開発時・CI |
| **Interaction Testing** | Storybook Interactions | ユーザー操作 | 開発時・CI |
| **Unit Testing** | Jest (utilities) | ビジネスロジック | 開発時・CI |
| **Lint/Type Check** | ESLint・TypeScript | コード品質 | 開発時・CI |
| **E2E Testing** | Manual | 重要フロー | リリース前 |

## 📖 Storybook中心の開発

### 設定ファイル

**.storybook/main.ts**:
```typescript
import type { StorybookConfig } from '@storybook/nextjs';

const config: StorybookConfig = {
  stories: ['../stories/**/*.stories.@(js|jsx|ts|tsx)'],
  addons: [
    '@storybook/addon-links',
    '@storybook/addon-essentials',      // Controls, Actions, Docs
    '@storybook/addon-interactions',    // インタラクションテスト
    'storybook-addon-mock',            // API モック
    'storybook-addon-jotai',           // Jotai 状態管理
  ],
  framework: {
    name: '@storybook/nextjs',
    options: {},
  },
  docs: {
    autodocs: 'tag',
  },
  staticDirs: ['../public'],
};

export default config;
```

### コンポーネントストーリー作成

```tsx
// stories/Auth/LoginForm.stories.tsx
import type { Meta, StoryObj } from '@storybook/react';
import { userEvent, within, expect } from '@storybook/test';
import { rest } from 'msw';
import { LoginForm } from '../../components/Login/LoginForm';

const meta: Meta<typeof LoginForm> = {
  title: 'Auth/LoginForm',
  component: LoginForm,
  parameters: {
    layout: 'centered',
    docs: {
      description: {
        component: 'メールアドレスとパスワードでログインするフォーム',
      },
    },
  },
  tags: ['autodocs'],
};

export default meta;
type Story = StoryObj<typeof meta>;

// 基本的なストーリー
export const Default: Story = {};

// ローディング状態
export const Loading: Story = {
  parameters: {
    msw: {
      handlers: [
        rest.post('/api/v2/login/password', (req, res, ctx) => {
          return res(ctx.delay('infinite')); // 無限ローディング
        }),
      ],
    },
  },
  play: async ({ canvasElement }) => {
    const canvas = within(canvasElement);
    const emailInput = canvas.getByPlaceholderText('メールアドレス');
    const passwordInput = canvas.getByPlaceholderText('パスワード');
    const submitButton = canvas.getByRole('button', { name: 'ログイン' });

    // フォーム入力
    await userEvent.type(emailInput, 'user@example.com');
    await userEvent.type(passwordInput, 'password123');
    await userEvent.click(submitButton);

    // ローディング状態確認
    await expect(canvas.getByText('ログイン中...')).toBeInTheDocument();
  },
};

// エラー状態
export const WithError: Story = {
  parameters: {
    msw: {
      handlers: [
        rest.post('/api/v2/login/password', (req, res, ctx) => {
          return res(
            ctx.status(401),
            ctx.json({ error: 'メールアドレスまたはパスワードが正しくありません' })
          );
        }),
      ],
    },
  },
  play: async ({ canvasElement }) => {
    const canvas = within(canvasElement);
    const emailInput = canvas.getByPlaceholderText('メールアドレス');
    const passwordInput = canvas.getByPlaceholderText('パスワード');
    const submitButton = canvas.getByRole('button', { name: 'ログイン' });

    // 無効な認証情報でログイン試行
    await userEvent.type(emailInput, 'invalid@example.com');
    await userEvent.type(passwordInput, 'wrongpassword');
    await userEvent.click(submitButton);

    // エラーメッセージ確認
    await expect(
      await canvas.findByText('メールアドレスまたはパスワードが正しくありません')
    ).toBeInTheDocument();
  },
};

// 成功フロー
export const Success: Story = {
  parameters: {
    msw: {
      handlers: [
        rest.post('/api/v2/login/password', (req, res, ctx) => {
          return res(
            ctx.status(200),
            ctx.json({
              user: { id: '1', email: 'user@example.com', userName: 'testuser' },
              token: 'mock-jwt-token'
            })
          );
        }),
      ],
    },
  },
  play: async ({ canvasElement }) => {
    const canvas = within(canvasElement);
    const emailInput = canvas.getByPlaceholderText('メールアドレス');
    const passwordInput = canvas.getByPlaceholderText('パスワード');
    const submitButton = canvas.getByRole('button', { name: 'ログイン' });

    // 有効な認証情報でログイン
    await userEvent.type(emailInput, 'user@example.com');
    await userEvent.type(passwordInput, 'correct-password');
    await userEvent.click(submitButton);

    // 成功メッセージ確認（トースト等）
    await expect(
      await canvas.findByText('ログインしました')
    ).toBeInTheDocument();
  },
};
```

### インタラクションテスト

```tsx
// stories/Client/ClientForm.stories.tsx
export const CreateClient: Story = {
  play: async ({ canvasElement }) => {
    const canvas = within(canvasElement);

    // フォーム入力
    const nameInput = canvas.getByLabelText('クライアント名');
    await userEvent.type(nameInput, 'テストアプリケーション');

    const descInput = canvas.getByLabelText('説明');
    await userEvent.type(descInput, 'テスト用のOIDCクライアント');

    const redirectInput = canvas.getByLabelText('リダイレクトURL');
    await userEvent.type(redirectInput, 'https://example.com/callback');

    // スコープ選択
    const openidScope = canvas.getByLabelText('openid');
    await userEvent.click(openidScope);

    const profileScope = canvas.getByLabelText('profile');
    await userEvent.click(profileScope);

    // フォーム送信
    const submitButton = canvas.getByRole('button', { name: '作成' });
    await userEvent.click(submitButton);

    // 成功確認
    await expect(
      await canvas.findByText('クライアントを作成しました')
    ).toBeInTheDocument();
  },
};
```

## 🎭 ビジュアルテスト

### コンポーネント状態の網羅

```tsx
// stories/Common/Button.stories.tsx
export const AllVariants: Story = {
  render: () => (
    <VStack spacing={4} align="stretch">
      <HStack spacing={4}>
        <Button variant="primary">Primary</Button>
        <Button variant="secondary">Secondary</Button>
        <Button variant="danger">Danger</Button>
      </HStack>

      <HStack spacing={4}>
        <Button size="sm">Small</Button>
        <Button size="md">Medium</Button>
        <Button size="lg">Large</Button>
      </HStack>

      <HStack spacing={4}>
        <Button isLoading>Loading</Button>
        <Button isDisabled>Disabled</Button>
      </HStack>

      <HStack spacing={4}>
        <Button leftIcon={<Icon as={FiUser} />}>With Icon</Button>
        <IconButton aria-label="Settings" icon={<Icon as={FiSettings} />} />
      </HStack>
    </VStack>
  ),
};

// レスポンシブ表示確認
export const Responsive: Story = {
  parameters: {
    viewport: {
      viewports: {
        mobile: { name: 'Mobile', styles: { width: '375px', height: '667px' } },
        tablet: { name: 'Tablet', styles: { width: '768px', height: '1024px' } },
        desktop: { name: 'Desktop', styles: { width: '1200px', height: '800px' } },
      },
    },
  },
  render: () => <ResponsiveComponent />,
};
```

### ダークモードテスト

```tsx
// stories/Common/Card.stories.tsx
export const DarkMode: Story = {
  parameters: {
    backgrounds: { default: 'dark' },
    chakra: { colorMode: 'dark' },
  },
  render: () => (
    <Card>
      <CardHeader>
        <Heading size="md">ダークモード</Heading>
      </CardHeader>
      <CardBody>
        <Text>ダークモードでの表示確認</Text>
      </CardBody>
    </Card>
  ),
};
```

## 🔌 モックとテストデータ

### MSW (Mock Service Worker) 活用

```tsx
// .storybook/preview.ts
import { initialize, mswDecorator } from 'msw-storybook-addon';

// MSW初期化
initialize();

export const decorators = [mswDecorator];

// グローバルパラメータ
export const parameters = {
  msw: {
    handlers: [
      // デフォルトのAPIモック
      rest.get('/api/v2/user/me', (req, res, ctx) => {
        return res(
          ctx.json({
            id: '1',
            email: 'user@example.com',
            userName: 'デモユーザー',
            avatarUrl: 'https://api.dicebear.com/7.x/personas/svg?seed=demo',
          })
        );
      }),
    ],
  },
};
```

### テストデータファクトリー

```tsx
// stories/utils/factories.ts
import { faker } from '@faker-js/faker/locale/ja';

export const createMockUser = (overrides: Partial<User> = {}): User => ({
  id: faker.string.ulid(),
  email: faker.internet.email(),
  userName: faker.internet.userName(),
  avatarUrl: faker.image.avatar(),
  createdAt: faker.date.past().toISOString(),
  updatedAt: faker.date.recent().toISOString(),
  ...overrides,
});

export const createMockClient = (overrides: Partial<OIDCClient> = {}): OIDCClient => ({
  id: faker.string.ulid(),
  name: faker.company.name() + 'アプリ',
  clientId: faker.string.uuid(),
  description: faker.lorem.sentence(),
  redirectUris: [faker.internet.url() + '/callback'],
  scopes: ['openid', 'profile', 'email'],
  createdAt: faker.date.past().toISOString(),
  ...overrides,
});

// 使用例
export const WithManyClients: Story = {
  parameters: {
    msw: {
      handlers: [
        rest.get('/api/v2/clients', (req, res, ctx) => {
          const clients = Array.from({ length: 10 }, () => createMockClient());
          return res(ctx.json(clients));
        }),
      ],
    },
  },
};
```

## 🎯 フォームテスト

### バリデーションテスト

```tsx
// stories/Forms/RegisterForm.stories.tsx
export const ValidationErrors: Story = {
  play: async ({ canvasElement }) => {
    const canvas = within(canvasElement);

    // 空のフォーム送信
    const submitButton = canvas.getByRole('button', { name: '登録' });
    await userEvent.click(submitButton);

    // バリデーションエラー確認
    await expect(canvas.getByText('メールアドレスは必須です')).toBeInTheDocument();
    await expect(canvas.getByText('ユーザー名は必須です')).toBeInTheDocument();

    // 無効な形式でのテスト
    const emailInput = canvas.getByLabelText('メールアドレス');
    await userEvent.type(emailInput, 'invalid-email');
    await userEvent.tab(); // フォーカス外し

    await expect(
      canvas.getByText('有効なメールアドレスを入力してください')
    ).toBeInTheDocument();

    // 文字数制限テスト
    const userNameInput = canvas.getByLabelText('ユーザー名');
    const longUserName = 'a'.repeat(51); // 50文字制限を超過
    await userEvent.type(userNameInput, longUserName);
    await userEvent.tab();

    await expect(
      canvas.getByText('ユーザー名は50文字以内で入力してください')
    ).toBeInTheDocument();
  },
};
```

## 📊 アクセシビリティテスト

### Storybookアクセシビリティ統合

```tsx
// .storybook/main.ts
export default {
  addons: [
    '@storybook/addon-a11y', // アクセシビリティチェック
  ],
};

// stories での設定
export const AccessibilityTest: Story = {
  parameters: {
    a11y: {
      config: {
        rules: [
          {
            id: 'color-contrast',
            enabled: true,
          },
          {
            id: 'focus-trap',
            enabled: true,
          },
        ],
      },
    },
  },
};
```

### キーボードナビゲーションテスト

```tsx
export const KeyboardNavigation: Story = {
  play: async ({ canvasElement }) => {
    const canvas = within(canvasElement);

    // Tab キーでのフォーカス遷移テスト
    await userEvent.tab();
    expect(canvas.getByLabelText('メールアドレス')).toHaveFocus();

    await userEvent.tab();
    expect(canvas.getByLabelText('パスワード')).toHaveFocus();

    await userEvent.tab();
    expect(canvas.getByRole('button', { name: 'ログイン' })).toHaveFocus();

    // Enter キーでの送信テスト
    await userEvent.keyboard('{Enter}');

    // フォーム送信の確認
    await expect(canvas.getByText('ログイン中...')).toBeInTheDocument();
  },
};
```

## 🧪 ユーティリティテスト

### Jest単体テスト

```tsx
// utils/__tests__/validate.test.ts
import { validateEmail, validateUserName, validatePassword } from '../validate';

describe('validate', () => {
  describe('validateEmail', () => {
    test('有効なメールアドレス', () => {
      expect(validateEmail('user@example.com')).toBe(true);
      expect(validateEmail('test.email+tag@domain.co.jp')).toBe(true);
    });

    test('無効なメールアドレス', () => {
      expect(validateEmail('invalid-email')).toBe(false);
      expect(validateEmail('user@')).toBe(false);
      expect(validateEmail('@domain.com')).toBe(false);
    });
  });

  describe('validateUserName', () => {
    test('有効なユーザー名', () => {
      expect(validateUserName('username')).toBe(true);
      expect(validateUserName('user-name')).toBe(true);
      expect(validateUserName('user123')).toBe(true);
    });

    test('無効なユーザー名', () => {
      expect(validateUserName('')).toBe(false);
      expect(validateUserName('a'.repeat(51))).toBe(false); // 50文字超過
      expect(validateUserName('user name')).toBe(false); // スペース含む
      expect(validateUserName('ユーザー名')).toBe(false); // 日本語
    });
  });
});
```

## 📈 テストカバレッジと品質指標

### Storybookカバレッジ

```bash
# Storybook起動とカバレッジ測定
pnpm storybook

# インタラクションテスト実行
pnpm test-storybook

# ビジュアルリグレッションテスト
pnpm storybook:chromatic
```

### 品質指標

| 指標 | 目標 | 測定方法 |
|------|------|----------|
| コンポーネントカバレッジ | 95%+ | Storybook Story数 |
| アクセシビリティ | AA準拠 | a11y addon |
| TypeScript Coverage | 100% | tsc --noEmit |
| Lint Errors | 0 | ESLint |

## 🔄 CI/CD統合

### GitHub Actions設定

```yaml
# .github/workflows/frontend-test.yml
name: Frontend Tests

on: [push, pull_request]

jobs:
  test:
    runs-on: ubuntu-latest

    steps:
      - uses: actions/checkout@v3
      - uses: actions/setup-node@v3
        with:
          node-version: '18'

      - name: Install dependencies
        run: pnpm install

      - name: Type check
        run: pnpm tsc --noEmit

      - name: Lint
        run: pnpm lint

      - name: Build Storybook
        run: pnpm build-storybook

      - name: Run interaction tests
        run: pnpm test-storybook
```

## 📚 関連ドキュメント

- [フロントエンドアーキテクチャ](./architecture.md)
- [コンポーネント設計](./components.md)
- [状態管理](./state-management.md)