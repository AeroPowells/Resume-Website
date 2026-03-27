import { useResumeSection } from '../hooks/useResume'
import styles from './Section.module.css'

export default function Education() {
  const { data, loading, error } = useResumeSection('/api/resume/education')

  if (loading) return <p className={styles.state}>Loading...</p>
  if (error) return <p className={styles.state}>Error: {error}</p>

  return (
    <section className={styles.section}>
      <h2 className={styles.heading}>Education</h2>
      {data.map((edu, i) => (
        <div key={i} className={styles.item}>
          <div className={styles.itemHeader}>
            <div>
              <h3 className={styles.itemTitle}>{edu.degree} in {edu.field}</h3>
              <p className={styles.itemSubtitle}>{edu.institution}</p>
            </div>
            <span className={styles.dates}>{edu.start_date} – {edu.end_date}</span>
          </div>
        </div>
      ))}
    </section>
  )
}
