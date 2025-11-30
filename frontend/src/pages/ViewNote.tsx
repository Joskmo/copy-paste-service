import { useEffect, useState } from 'react';
import { useParams, Link } from 'react-router-dom';
import { getNote, type GetNoteResponse } from '../api';
import styles from './ViewNote.module.css';

export function ViewNote() {
  const { id } = useParams<{ id: string }>();
  const [note, setNote] = useState<GetNoteResponse | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [copied, setCopied] = useState(false);

  useEffect(() => {
    if (!id) return;

    const fetchNote = async () => {
      try {
        const data = await getNote(id);
        setNote(data);
      } catch (err) {
        setError(err instanceof Error ? err.message : 'Произошла ошибка');
      } finally {
        setLoading(false);
      }
    };

    fetchNote();
  }, [id]);

  const handleCopy = async () => {
    if (note) {
      await navigator.clipboard.writeText(note.content);
      setCopied(true);
      setTimeout(() => setCopied(false), 2000);
    }
  };

  if (loading) {
    return (
      <div className={styles.container}>
        <div className={styles.card}>
          <div className={styles.loading}>Загрузка...</div>
        </div>
      </div>
    );
  }

  if (error) {
    return (
      <div className={styles.container}>
        <div className={styles.card}>
          <div className={styles.errorIcon}>✕</div>
          <h2 className={styles.errorTitle}>Заметка не найдена</h2>
          <p className={styles.errorText}>{error}</p>
          <Link to="/" className={styles.backLink}>
            ← Создать новую заметку
          </Link>
        </div>
      </div>
    );
  }

  return (
    <div className={styles.container}>
      <div className={styles.card}>
        <div className={styles.header}>
          <span className={styles.noteId}>{id}</span>
          <button 
            className={styles.copyButton}
            onClick={handleCopy}
          >
            {copied ? '✓ Скопировано' : 'Копировать в буфер'}
          </button>
        </div>

        <div className={styles.contentBox}>
          <pre className={styles.content}>{note?.content}</pre>
        </div>

        <Link to="/" className={styles.newLink}>
          ← Создать новую заметку
        </Link>
      </div>
    </div>
  );
}

