<?php

namespace App\Http\Requests;

class ArticleRequest extends FormRequest
{
    /**
     * Get the validation rules that apply to the request.
     *
     * @return array<string, mixed>
     */
    public function rules()
    {
        return match ($this->getMethod()) {
            'POST', 'PATCH' => [
                'title' => 'required|min:2',
                'type' => 'in:markdown,body',
                'content.body' => 'required_if:type,body',
                'content.markdown' => 'required_if:type,markdown',
            ],
            default => [
                //
            ],
        };
    }
}
