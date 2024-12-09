# ミニマムTwitterアプリ設計

## 概要

### 機能一覧
1. **ユーザー登録・ログイン**
   - ユーザー名とパスワードを登録し、ログイン認証。
2. **投稿（ツイート）**
   - テキスト投稿と投稿時刻の記録。
3. **タイムライン**
   - 自分およびフォローしているユーザーの投稿を時系列順で取得。
4. **フォロー/フォロー解除**
   - 他のユーザーをフォローまたはフォロー解除。

---

## 設計

### ドメインモデル

#### エンティティ
- **User**
  - ID（不変）
  - ユーザー名
  - フォロー関係（フォロワーとフォローイー）

- **Tweet**
  - ID（不変）
  - 投稿者（User ID）
  - 投稿本文
  - 投稿日時

#### 値オブジェクト
- **UserName**
  - ユーザー名の制約（例: 非空、一意）
- **TweetContent**
  - ツイート本文の制約（例: 最大140文字）

---

### アプリケーションレイヤ（ユースケース）

#### ユースケース一覧
1. **ユーザー関連**
   - ユーザー登録（SignUp）
   - ログイン（SignIn）
   - フォロー/フォロー解除（Follow/Unfollow）
2. **ツイート関連**
   - ツイート投稿（PostTweet）
   - タイムライン取得（GetTimeline）

#### ユースケース例
- **ユーザー登録（SignUp）**
  - 入力: ユーザー名、パスワード
  - 出力: 新規ユーザーID
- **タイムライン取得（GetTimeline）**
  - 入力: ユーザーID
  - 出力: フォロー中のユーザー＋自身の投稿の時系列リスト

---

### インフラストラクチャーレイヤ

#### リポジトリ
- **UserRepository**
  - `Save(user User) error`
  - `FindByID(id string) (User, error)`
- **TweetRepository**
  - `Save(tweet Tweet) error`
  - `FindByUserID(userID string) ([]Tweet, error)`
  - `FindTimeline(userID string) ([]Tweet, error)`

#### HTTPハンドラ
- **エンドポイント例**
  - `POST /users` → ユーザー登録
  - `POST /login` → ログイン
  - `POST /tweets` → ツイート投稿
  - `GET /timeline` → タイムライン取得
  - `POST /users/{id}/follow` → フォロー/フォロー解除

---

### データベース設計

#### テーブル構成
1. **users**
   - `id` (PK)
   - `username`
   - `password_hash`

2. **tweets**
   - `id` (PK)
   - `content`
   - `author_id` (FK)
   - `timestamp`

3. **follows**
   - `follower_id` (FK)
   - `followee_id` (FK)

---