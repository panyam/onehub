import { MouseEvent, useEffect, useRef } from "react";
import styled from "styled-components";
import styles from '@/components/styles/DialogModal.module.css'

const Buttons = styled.div`
  display: flex;
  gap: 20px;
`;

const isClickInsideRectangle = (e: MouseEvent, element: HTMLElement) => {
  const r = element.getBoundingClientRect();

  return (
    e.clientX > r.left &&
    e.clientX < r.right &&
    e.clientY > r.top &&
    e.clientY < r.bottom
  );
};

type Props = {
  title: string;
  children: React.ReactNode;
  isOpened: boolean;
  onClose?: (label: String) => boolean;
  buttons?: string[];
  onCancel?: () => boolean;
  cancelButtonLabel?: string;
};

const DialogModal = ({
  title,
  isOpened,
  onCancel,
  onClose,
  buttons,
  children,
  cancelButtonLabel,
}: Props) => {
  const ref = useRef<HTMLDialogElement>(null);

  useEffect(() => {
    if (isOpened) {
      ref.current?.showModal();
      document.body.classList.add("modal-open"); // prevent bg scroll
    } else {
      ref.current?.close();
      document.body.classList.remove("modal-open");
    }
  }, [isOpened]);

  buttons = buttons || ["Ok"]
  cancelButtonLabel = cancelButtonLabel || "Cancel"
  if (!onCancel || onCancel == null) {
    onCancel = () => true
  }
  if (!onClose || onClose == null) {
    onClose = (label) => true
  }

  return (
  <div className = {styles.dialog_container}>
    <dialog className = {styles.dialog}
      ref={ref}
      onCancel={onCancel}
      onClick={(e) =>
        ref.current && !isClickInsideRectangle(e, ref.current) && onCancel!()
      }
    >
      <h3>{title}</h3>

      {children}

      <Buttons>
        <center>
          {
            buttons.map((label, index) => 
              <button key={index}
                      className={styles.dialog_button}
                      onClick={() => onClose!(label)}>{label}
                  </button>
            )
          }
          <button key="cancel" className={styles.dialog_button}
                  onClick={() => onCancel!()}>{cancelButtonLabel}</button>
        </center>
      </Buttons>
    </dialog>
  </div>
  );
};

export default DialogModal;
