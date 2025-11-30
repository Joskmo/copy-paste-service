import { useState } from 'react';
import { createNote } from '../api';
import { APP_URL } from '../config';
import styles from './CreateNote.module.css';

export function CreateNote() {
  const [content, setContent] = useState('');
  const [noteUrl, setNoteUrl] = useState<string | null>(null);
  const [expiresAt, setExpiresAt] = useState<string | null>(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [copied, setCopied] = useState(false);

  const handleSubmit = async () => {
    if (!content.trim()) {
      setError('Введите текст');
      return;
    }

    setLoading(true);
    setError(null);

    try {
      const response = await createNote(content);
      // Формируем ссылку на фронтенд, а не на API
      const frontendUrl = `${APP_URL}/${response.id}`;
      setNoteUrl(frontendUrl);
      setExpiresAt(response.expires_at);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Произошла ошибка');
    } finally {
      setLoading(false);
    }
  };

  const handleCopyLink = async () => {
    if (noteUrl) {
      await navigator.clipboard.writeText(noteUrl);
      setCopied(true);
      setTimeout(() => setCopied(false), 2000);
    }
  };

  const handleReset = () => {
    setContent('');
    setNoteUrl(null);
    setExpiresAt(null);
    setError(null);
  };

  const formatExpiresAt = (dateString: string) => {
    const date = new Date(dateString);
    return date.toLocaleString('ru-RU', {
      day: 'numeric',
      month: 'long',
      hour: '2-digit',
      minute: '2-digit',
    });
  };

  if (noteUrl) {
    return (
      <div className={styles.container}>
        <div className={styles.card}>
          <div className={styles.successIcon}>✓</div>
          <h2 className={styles.successTitle}>Ссылка готова!</h2>
          
          <div className={styles.linkBox}>
            <span className={styles.link}>{noteUrl}</span>
            <button 
              className={styles.copyButton}
              onClick={handleCopyLink}
            >
              {copied ? '✓ Скопировано' : 'Копировать'}
            </button>
          </div>

          <p className={styles.expiresText}>
            Ссылка действительна до {expiresAt && formatExpiresAt(expiresAt)}
          </p>

          <button className={styles.newButton} onClick={handleReset}>
            Создать новую заметку
          </button>
        </div>
      </div>
    );
  }

  return (
    <div className={styles.container}>
      <div className={styles.card}>
        <h1 className={styles.title}>Быстрый обмен текстом</h1>
        <p className={styles.subtitle}>
          Вставьте текст и получите короткую ссылку. Хранится 3 часа.
        </p>

        <textarea
          className={styles.textarea}
          placeholder="Вставьте ваш текст здесь..."
          value={content}
          onChange={(e) => setContent(e.target.value)}
          rows={10}
        />

        {error && <p className={styles.error}>{error}</p>}

        <button
          className={styles.submitButton}
          onClick={handleSubmit}
          disabled={loading}
        >
          {loading ? 'Создание...' : 'Получить ссылку'}
        </button>
      </div>
    </div>
  );
}

