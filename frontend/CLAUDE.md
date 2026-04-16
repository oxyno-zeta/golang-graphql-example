# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Instructions for CLAUDE

- If you see something missing in following sections, add it
- If you see that something have been deleted in project but not in this file, inform and propose a deletion
- If you see that something needs to be updated, update it
- When you are changing, adding or deleting code, run `yarn lint` to ensure linter is passing

## Commands

All commands are run from the `frontend/` directory.

```bash
yarn start             # Dev server at http://localhost:3000 (proxies /api → localhost:8080)
yarn build             # Production build (output: dist/)
yarn preview           # Preview production build locally
yarn test              # Jest tests
yarn test:update       # Vitest — update snapshots
yarn test:coverage     # Jest tests with coverage report (output: coverage/)
yarn lint              # ESLint with auto-fix
yarn lint:tsc          # TypeScript type check only
yarn prettier          # Format with Prettier
yarn storybook         # Storybook dev server at http://localhost:6006
yarn build-storybook   # Build Storybook (output: storybook-static/)
```

Run a single test:

```bash
yarn test --testPathPattern=ComponentName
```

## Architecture

### Tech Stack

| Layer     | Library                  |
| --------- | ------------------------ |
| Framework | React                    |
| Build     | Vite                     |
| UI        | MUI (`@mui/material`)    |
| Data Grid | `@mui/x-data-grid`       |
| GraphQL   | Apollo Client            |
| Forms     | react-hook-form + Yup    |
| Routing   | react-router             |
| i18n      | i18next + react-i18next  |
| Date/time | dayjs                    |
| Testing   | Vitest + Testing Library |
| Storybook | Storybook (Vite builder) |

### Directory Layout

```
src/
├── components/          # Reusable UI components (1 component per file, 1 file per folder)
├── contexts/            # React Context providers
├── models/              # TypeScript interfaces/types
├── routes/              # Route definitions and page components
├── utils/               # Utility functions
├── App.tsx              # Root layout (ThemeProvider, RouterOutlet)
├── index.tsx            # Entry point
├── i18n.ts              # i18n setup
└── yup-i18n.ts          # Yup validation message i18n
public/
├── config/config.json   # Runtime config (graphqlEndpoint, configCookieDomain)
└── i18n-{lng}.json      # Translation files per language
```

### Component File Structure

Each component lives in its own folder — **one component per file, one file per folder**:

```
ComponentName/
├── ComponentName.tsx          # Implementation
├── ComponentName.test.tsx     # Jest tests
├── ComponentName.stories.tsx  # Storybook stories
└── index.tsx                  # Barrel export: export { default } from './ComponentName'
```

Sub-components needed only by a parent live in a nested `components/` subfolder inside the parent folder.

### Routing

Defined in `src/routes/router-routes.tsx` using `createBrowserRouter()`:

| Path                 | Component              |
| -------------------- | ---------------------- |
| `/`                  | Index                  |
| `/redirect-to/:name` | GraphQL-based redirect |
| `*`                  | 404 Not Found          |

URL search params encode all filter/sort/pagination state (bookmarkable). Use the `useJSONObjectFromSearchParam()` utility for this.

### State Management

No Redux or Zustand. State is handled by:

- **Apollo Client** — all server data (queries, mutations, cache)
- **React Context** — cross-cutting UI state (theme, config, drawer, timezone)
- **react-hook-form** — local form state
- **URL search params** — filter/sort/pagination state

### GraphQL / Apollo Client

Configured in `src/components/ClientProvider/ClientProvider.tsx`.

- Endpoint: from runtime config (`cfg.graphqlEndpoint`, default `/api/graphql`)
- Credentials: `'include'` (cookie-based OAuth2 session)
- Link chain: ErrorLink (401 reload) → force Accept header → trace error link → HttpLink
- Cache: InMemoryCache, no persistence

### i18n

- Translations loaded from `/i18n-{lng}.json` at runtime (HTTP backend)
- Browser language auto-detected, fallback: `en`
- All user-visible strings must use `useTranslation()` — no hardcoded UI text
- Yup validation messages are also i18n-keyed via `src/yup-i18n.ts`

### Configuration

Runtime config loaded via axios from `/public/config/config.json`:

```json
{
  "graphqlEndpoint": "/api/graphql",
  "configCookieDomain": "localhost"
}
```

Access it via `useContext(ConfigContext)` — do not use env vars for runtime values.

**TypeScript path aliases** (use these for imports):

| Alias           | Resolves to        |
| --------------- | ------------------ |
| `~components/*` | `src/components/*` |
| `~contexts/*`   | `src/contexts/*`   |
| `~models/*`     | `src/models/*`     |
| `~routes/*`     | `src/routes/*`     |
| `~utils/*`      | `src/utils/*`      |

### Theming

- `ThemeProvider` in `src/components/theming/ThemeProvider/` wraps the entire app
- Light/dark mode toggled via `ColorModeContext` (stored in a cookie, falls back to system preference)
- Theme customization: pass `themeOptions` to `ThemeProvider` — do not call `createTheme()` elsewhere
- `CssBaseline` is applied once at the app root

## Development Guidelines

### Use MUI at Maximum

Always prefer MUI components over custom HTML or third-party UI libraries:

- **Layout:** `Box`, `Stack`, `Grid2`, `Container`
- **Typography:** `Typography` (never raw `<p>`, `<h1>`, etc.)
- **Forms:** `TextField`, `Select`, `Checkbox`, `Switch`, `Autocomplete` — wrapped with react-hook-form `Controller`
- **Feedback:** `CircularProgress`, `Alert`, `Snackbar`, `Skeleton`
- **Navigation:** `AppBar`, `Drawer`, `Tabs`, `Breadcrumbs`
- **Data:** `DataGrid` from `@mui/x-data-grid`
- **Icons:** `@mdi/js` via a thin SvgIcon wrapper — do not import from `@mui/icons-material`
- **Dates:** `@mui/x-date-pickers` with dayjs adapter

**Styling:** use the `sx` prop for all component styles. No separate CSS files, no `styled()` calls unless `sx` is genuinely insufficient. But prefer inline `style={{}}` if possible.

### One Component Per File

- Each file exports exactly one React component.
- Keep the component name and file name identical.
- Co-locate tests (`.test.tsx`) and stories (`.stories.tsx`) in the same folder.
- If a component grows complex, extract sub-components into a `components/` subfolder — do not put them in the same file.

### Forms

- Always use **react-hook-form** with `useForm()` and `Controller`.
- Validation schemas defined with **Yup** — all messages must be i18n keys.
- Use the existing form wrappers (`FormInput`, `FormAutocomplete`, etc.) in `src/components/form/` before creating new ones.

### Testing

- **Runner:** Vitest — config in `vitest.config.mts`, setup in `.vitest/vitest.setup.ts`.
- Tests use `@testing-library/react` with role queries (accessibility-first).
- Use `vi.*` globals (`vi.fn()`, `vi.mock()`, `vi.spyOn()`) — do not use `jest.*`.
- Mock i18n in tests: `vi.mock('react-i18next', () => ({ useTranslation: () => ({ t: (key) => key }) }))`
- Snapshot tests are acceptable for pure presentational components. Update with `yarn test:update`.
- No mocking of Apollo Client — use `MockedProvider` from `@apollo/client/testing`.
- `@testing-library/jest-dom` matchers (`toHaveClass`, `toHaveTextContent`, etc.) are available globally via the setup file.

### Storybook

Every reusable component should have a `.stories.tsx` file. The global Storybook decorator applies `ThemeProvider` automatically, so stories don't need to wrap components manually.

### Code Style

- **Prettier:** 120-char line width, 2-space indent, single quotes, trailing commas, arrow parens always.
- **Imports:** use path aliases (`~components/`, `~utils/`, etc.) — no relative `../../` imports across src boundaries.
- **No console.log** — enforced by ESLint.
- **No unused imports** — enforced by `eslint-plugin-unused-imports`.

## Commit Convention

Follow Angular commit convention: `feat:`, `fix:`, `docs:`, `chore:`, `refactor:`, etc. Enforced by pre-commit hooks via Commitizen.
