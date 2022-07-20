<?php

namespace App\Mail;

use  Illuminate\Support\Str;
use Illuminate\Bus\Queueable;
use Illuminate\Contracts\Queue\ShouldQueue;
use Illuminate\Mail\Mailable;
use Illuminate\Queue\SerializesModels;

class ResetPassword extends Mailable implements ShouldQueue
{
    use Queueable, SerializesModels;

    public string $email;

    public string $token;

    public string $link;

    /**
     * Create a new message instance.
     *
     * @param string $email
     * @param string $token
     * @return void
     */
    public function __construct(string $email, string $token)
    {
        $this->email = $email;
        $this->token = $token;

        $this->link = $this->resetPasswordLink();
    }

    public function resetPasswordLink(): string
    {
        $params = [
            'token' => $this->token,
            'email' => $this->email,
        ];

        return Str::finish(config('app.site_url'), '/').'reset-password?'.http_build_query($params);
    }

    /**
     * Build the message.
     *
     * @return $this
     */
    public function build(): static
    {
        return $this->subject('重置密码')->markdown('mails.reset-password');
    }
}
